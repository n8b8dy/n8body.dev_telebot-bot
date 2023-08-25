package handlers

import (
	"gopkg.in/telebot.v3"
	"n8body.dev/telebot-bot/app/controllers"
)

func EntertainingHandlers(bot *telebot.Bot, controller *controllers.MainController) {
	entertaining := bot.Group()
	entertaining.Handle("/meme", controller.MemeCommand)
	entertaining.Handle("/memes", controller.MemeCommand)
}
