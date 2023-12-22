package handlers

import (
	"fmt"

	"github.com/UjjwalMahar/captain-ai/initializers"
	"github.com/UjjwalMahar/captain-ai/ai"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var UserInput string
var Msg *tgbotapi.MessageConfig

func HandelTelegramBot(){
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 80
	bot := initializers.TeleBotInitializer()
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			Msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			UserInput := update.Message.Text
		
		switch UserInput {
		case "/start":
			Msg.Text = "Hello this is Captain AI ⚓\nHow may I help you ?"
			bot.Send(Msg)
		case "/help":
			Msg.Text = "Just type what you want to ask !"
			bot.Send(Msg)
		default:
			Msg.Text, _ = ai.GenerateAiContent(UserInput)
			fmt.Print(ai.GenerateAiContent(UserInput))
			bot.Send(Msg)
		if _, err := fmt.Printf("All good"); err != nil {
			fmt.Print("The error is ",err)
		}
	}
	
	}
}

}