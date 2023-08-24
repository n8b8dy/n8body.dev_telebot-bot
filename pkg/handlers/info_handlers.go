package handlers

import (
	"gopkg.in/telebot.v3"
	"n8body.dev/telebot-bot/app/controllers"
)

func InfoHandlers(bot *telebot.Bot, controller *controllers.MainController) {
	info := bot.Group()
	info.Handle("/start", controller.StartCommand)
	info.Handle("/help", controller.HelpCommand)
	info.Handle("/commands", controller.HelpCommand)
	info.Handle("/me", controller.MeCommand)
}
