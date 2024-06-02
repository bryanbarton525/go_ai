package gemini

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	genai "github.com/google/generative-ai-go/genai"
)

const (
	// JSONDataFile represents the file path to the JSON data file
	JSONDataFile    = "dataset_k8s.json"
	OldJsonDataFile = "dataset_k8s.json.bak"
)

func Run(model *genai.GenerativeModel, ctx context.Context) error {
	// Read user input from command line
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter your summerization request: ")
	scanner.Scan()
	input := scanner.Text()
	// 1. Load JSON Data from File
	jsonData, err := os.ReadFile("dataset_k8s.json") // Load from file
	if err != nil {
		return err
	}
	err = os.WriteFile("dataset_k8s.json.bak", jsonData, 0644) // Backup the file
	if err != nil {
		return err
	}

	// 2. Generate Summary (Separate Function)
	reqResponse, err := handleRequest(input, ctx, model, jsonData)
	if err != nil {
		return errors.New("Failed to generate summary: " + err.Error())
	}

	// 3. Write Summary to File
	err = os.WriteFile("dataset_k8s.json", reqResponse, 0644)
	if err != nil {
		return errors.New("Failed to write summary to file: " + err.Error())
	}

	fmt.Println("Summary: ", string(reqResponse))

	// 3. Return Summary as JSON
	return nil
}

// handleRequest handles the incoming request by generating updated and properly formatted JSON based on the given customer support ticket data.
// It takes the input string, context, generative model, and JSON data as parameters.
// It returns the updated JSON data and an error if any.
func handleRequest(input string, ctx context.Context, model *genai.GenerativeModel, jsonData []byte) ([]byte, error) {
	var Candidates Candidates
	var tickets []Ticket
	prompt := fmt.Sprintf("Given the following customer support ticket data, %s and return the updated and properly formmated json:\n\n%s", input, string(jsonData))
	fmt.Println(prompt)
	response, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("generate content error: %v", err)
	}
	resp, err := json.Marshal(response)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %v", err)
	}
	err = json.Unmarshal(resp, &Candidates)
	if err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}
	for _, candidate := range Candidates.Candidates {
		tickets, err = extractAndConvertJSON(candidate.Content.Parts)
		fmt.Println("Tickets: ", tickets)
		if err != nil {
			return nil, fmt.Errorf("extract and convert JSON error: %v", err)
		}
	}
	updatedJson, err := json.Marshal(tickets)
	if err != nil {
		return nil, fmt.Errorf("json marshal error: %v", err)
	}

	fmt.Println("Updated Json: ", string(updatedJson))
	return updatedJson, nil
}

func extractAndConvertJSON(parts []string) ([]Ticket, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("no parts found in response")
	}

	// Get the raw JSON string (assuming it's the first part)
	rawJSON := parts[0]

	// Sanitize the JSON string (adjust as needed)
	sanitizedJSON := strings.TrimSpace(rawJSON)
	sanitizedJSON = strings.TrimPrefix(sanitizedJSON, "```json")
	sanitizedJSON = strings.TrimSuffix(sanitizedJSON, "```")

	// Unmarshal the JSON into a slice of Ticket structs
	var tickets []Ticket
	if err := json.Unmarshal([]byte(sanitizedJSON), &tickets); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	return tickets, nil
}
