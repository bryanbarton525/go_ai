package gemini

import (
	"context"
	"fmt"
	"log"

	genai "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
)

func Stream(model *genai.GenerativeModel, ctx context.Context) {
	// Streaming

	// Create a stream
	stream := model.GenerateContentStream(ctx, genai.Text("what is the sum of 37 and 64?"))

	// Read from the stream
	for {
		resp, err := stream.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Print the generated text
		fmt.Println(resp.Candidates[0].Content.Parts[0])
		fmt.Println(resp.UsageMetadata.CandidatesTokenCount)
	}
}
