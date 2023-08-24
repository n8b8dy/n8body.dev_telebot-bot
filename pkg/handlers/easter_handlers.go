package handlers

import (
	"gopkg.in/telebot.v3"
	"n8body.dev/telebot-bot/app/controllers"
)

func EasterHandlers(bot *telebot.Bot, controller *controllers.MainController) {
	easter := bot.Group()
	easter.Handle(telebot.OnText, controller.DaPizda)
}
