package initializers

//This is for initializing the telegram-bot
import (
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)
var Bot *tgbotapi.BotAPI

func TeleBotInitializer() *tgbotapi.BotAPI{
	Bot , err := tgbotapi.NewBotAPI((os.Getenv("TELEGRAM_TOKEN")))
	if err != nil{
		fmt.Println("Error loading telegram-bot token ", err)
	}
	Bot.Debug = true
	return Bot
}