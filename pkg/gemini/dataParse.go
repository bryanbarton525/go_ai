package gemini

import (
	"bufio"
	"context"
	"errors"
	"fmt" // Added for reading JSON file
	"os"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
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

	// 2. Generate Summary (Separate Function)
	err = generateSummary(input, ctx, model, jsonData)
	if err != nil {
		return errors.New("Failed to generate summary: " + err.Error())
	}
	return nil
}

// Function to generate the summary from JSON data
func generateSummary(input string, ctx context.Context, model *genai.GenerativeModel, jsonData []byte) error {
	prompt := fmt.Sprintf("Given the following customer support ticket data, %s:\n\n%s", input, string(jsonData))

	stream := model.GenerateContentStream(ctx, genai.Text(prompt))
	for {
		response, err := stream.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("stream error: %v", err)
		}
		for _, candidate := range response.Candidates {
			if candidate.Content.Parts[0] != nil {
				fmt.Println(candidate.Content.Parts[0])
			}
		}
	}
	return nil
}
