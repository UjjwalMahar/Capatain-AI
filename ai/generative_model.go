package ai

import (
	"fmt"
	"log"

	"github.com/UjjwalMahar/captain-ai/initializers"
	"github.com/google/generative-ai-go/genai"
)

var GenResponse string

//NOTE: EnsureGeminiModel initializes the Gemini model and ensures it's ready to use.
func EnsureGeminiModel() {
	client, err := initializers.GeminiInitializer()
	if err != nil {
		log.Fatal(err)
	}
	initializers.Client = client
}

func GenerateAiContent(input string) (string, error) {
	// Ensure the Gemini model is initialized
	EnsureGeminiModel()

	model := initializers.Client.GenerativeModel("gemini-pro")

	if model == nil {
		return "", fmt.Errorf("generative model is nil")
	}

	resp, err := model.GenerateContent(initializers.Ctx, genai.Text(input))
	if err != nil {
		return "", fmt.Errorf("error generating the response: %w", err)
	}

	if resp != nil {
		for _, cand := range resp.Candidates {
			for _, part := range cand.Content.Parts {
				GenResponse = fmt.Sprint(part)
				// fmt.Print(GenResponse)
			}
		}
	}

	return GenResponse, nil
}
