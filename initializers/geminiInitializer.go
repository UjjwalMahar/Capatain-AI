package initializers

import (
	"context"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var Ctx context.Context
var Client *genai.Client

// GeminiInitializer initializes the gemini model and returns the client.
func GeminiInitializer() (*genai.Client, error) {
	Ctx = context.Background()
	client, err := genai.NewClient(Ctx, option.WithAPIKey(os.Getenv("AI_API_KEY")))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return client, nil
}
