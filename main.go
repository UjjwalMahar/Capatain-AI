package main

import (
	"log"

	"github.com/UjjwalMahar/captain-ai/handlers"
	"github.com/UjjwalMahar/captain-ai/initializers"
)

func init(){
	initializers.LoadEnv()
	initializers.GeminiInitializer()
	initializers.TeleBotInitializer()
}

func main() {
	handlers.HandelTelegramBot()
	client, err := initializers.GeminiInitializer()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	
	}

