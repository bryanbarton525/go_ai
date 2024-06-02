package main

import (
	"log"

	"github.com/bryanbarton525/go_ai/pkg/gemini"
)

func main() {

	model, ctx, err := gemini.SetupGeminiClient()
	if err != nil {
		log.Fatalf("Failed to set up Gemini client: %v", err)
	}

	gemini.Run(model, ctx)
}
