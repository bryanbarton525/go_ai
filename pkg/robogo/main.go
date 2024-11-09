package robogo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/bryanbarton525/go_ai/pkg/gemini"
	genai "github.com/google/generative-ai-go/genai"
	"gocv.io/x/gocv"
)

func ControlRobot() {
	// Initialize Gemini client
	model, ctx, err := gemini.SetupGeminiClient()
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

	// ... (Preprocess image: resize, format conversion, etc.) ...

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
	// prompt := "Based on this image, which direction should the robot move? (forward, backward, left, right)"
	prompt := "Based on this image, describe to me what you see"
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

	// ... (Parse resp.Candidates[0] to extract direction) ...
	direction := resp.Candidates[0].Content.Parts[0]

	return direction, nil
}

func moveRobot(direction string) {
	// ... (Translate direction into robot commands) ...
	fmt.Printf("Moving robot: %s\n", direction)
}
