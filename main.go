package main

import (
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"
	"n8body.dev/telebot-bot/app/controllers"
	"n8body.dev/telebot-bot/pkg/handlers"
	"n8body.dev/telebot-bot/platform/database"
)

func main() {
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  os.Getenv("BOT_TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 8 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connecting to database...")
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done!")

	controller := controllers.NewMainController(db)

	bot.Use(middleware.Logger())
	bot.Use(middleware.Recover())

	if err := bot.SetCommands([]telebot.Command{
		{"/start", "shows initial message"},
		{"/help", "shows the list of all bot functions"},
		{"/commands", "alias for /help"},
		{"/bot", "shows story about the bot (AI powered, xd)"},
		{"/me", "shows information about the user"},
		{"/chat", "shows information about the chat (cannot be used in private chats)"},
		{"/meme", "sends a random meme"},
	}); err != nil {
		log.Fatal(err)
	}

	handlers.InfoHandlers(bot, controller)
	handlers.StickerHandlers(bot, controller)
	handlers.EntertainingHandlers(bot, controller)
	handlers.EasterHandlers(bot, controller)

	log.Println("Starting the bot...")
	bot.Start()
}
