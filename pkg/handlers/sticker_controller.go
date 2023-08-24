package handlers

import (
	"gopkg.in/telebot.v3"
	"n8body.dev/telebot-bot/app/controllers"
)

func StickerHandlers(bot *telebot.Bot, controller *controllers.MainController) {
	sticker := bot.Group()
	sticker.Handle(telebot.OnSticker, controller.Sticker)
	sticker.Handle("/sticker", controller.StickerCommand)
}
