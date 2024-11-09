package robogo

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/bryanbarton525/go_ai/pkg/gemini"
	genai "github.com/google/generative-ai-go/genai"
	"gocv.io/x/gocv"
)

var (
	model                  *genai.GenerativeModel
	ctx                    context.Context
	movementHistory        []string
	environmentDescription string
	prompt                 string // Declare prompt outside the function
)

func ControlRobot() {
	var err error
	// Initialize Gemini client
	model, ctx, err = gemini.SetupGeminiClient()
	if err != nil {
		log.Fatalf("Failed to set up Gemini client: %v", err)
	}

	// Initialize camera
	webcam, err := gocv.VideoCaptureDevice(0) // Assuming camera index 0
	if err != nil {
		log.Fatalf("Failed to initialize camera: %v", err)
	}
	defer webcam.Close()

	for {
		img, err := captureAndProcessImage(webcam)
		if err != nil {
			log.Printf("Failed to capture and process image: %v", err)
			return
		}
		direction, err := getDirectionFromGemini(model, img, ctx)
		if err != nil {
			log.Printf("Failed to get direction from Gemini: %v", err)
			continue
		}
		// moveRobot(direction.(string))
		fmt.Println(direction)
		time.Sleep(10 * time.Second)
	}
}

func captureAndProcessImage(webcam *gocv.VideoCapture) ([]byte, error) {
	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		return nil, fmt.Errorf("failed to read frame from camera")
	}

	// Encode the image as JPEG
	buf, err := gocv.IMEncode(".jpg", img)
	if err != nil {
		return nil, fmt.Errorf("failed to encode image: %w", err)
	}
	defer buf.Close()

	// Return the image data as []byte
	return buf.GetBytes(), nil
}

func getDirectionFromGemini(model *genai.GenerativeModel, imgData []byte, ctx context.Context) (interface{}, error) {
	prompt = fmt.Sprintf(
		"The robot has made the following moves: %s. "+
			"Its current understanding of the environment is: %s. "+
			"Based on this image, which direction should the robot move next "+
			"to explore the entire environment for image mapping? "+
			"Please format response to the following: 'direction (right, left, forward, backward) | description of image' The movement responses map to the following directions: "+
			"forward: moves 2 feet forward, backward: rotates robot 180 degrees, left: rotates 90 degress to the left, right: rotates 90 degrees to the right.",
		movementHistory, environmentDescription,
	)
	resp, err := model.GenerateContent(
		ctx,
		genai.Text(prompt),
		genai.ImageData("jpeg", imgData), // Use ImageData with the byte slice
	)
	if err != nil {
		return "", fmt.Errorf("failed to generate content: %w", err)
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no candidates found in response")
	}

	suggestion := resp.Candidates[0].Content.Parts[0]
	movement := strings.Split(convertPartToString(suggestion), " | ")[0]
	environmentDescription = strings.Split(convertPartToString(suggestion), " | ")[1]
	movementHistory = append(movementHistory, movement)

	return suggestion, nil
}

func moveRobot(direction string) {
	//TODO Implement robot movement based on direction. Use Firmata and Gobot.
	fmt.Printf("Moving robot: %s\n", direction)
}

func convertPartToString(part genai.Part) string {
	return string(part.(genai.Text))
}
