package main

import (
	"ai-test/internal/gemini"
	"log"
)

func main() {

	model, ctx, err := gemini.SetupGeminiClient()
	if err != nil {
		log.Fatalf("Failed to set up Gemini client: %v", err)
	}

	gemini.Run(model, ctx)
}
