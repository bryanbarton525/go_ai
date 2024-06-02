package gemini

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func SetupGeminiClient() (*genai.GenerativeModel, context.Context, error) {
	// Load API key from .env file (replace with your actual path)
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	apiKey := os.Getenv("GEMINI_API_KEY")

	// Initialize Gemini client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal("Error creating client:", err)
	}
	// defer client.Close()

	// Prepare the text generation request
	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")

	return model, ctx, nil

}
