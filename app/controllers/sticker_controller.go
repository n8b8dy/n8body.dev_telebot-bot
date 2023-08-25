package controllers

import (
	"fmt"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
)

type StickerController struct {
	*gorm.DB
}

func (*StickerController) Sticker(ctx telebot.Context) error {
	chat := ctx.Chat()
	if chat.Type != telebot.ChatPrivate {
		return nil
	}

	message := ctx.Message()
	stickerID := message.Sticker.File.FileID

	response := fmt.Sprintf("Sticker ID: <code>%s</code>", stickerID)

	return ctx.Reply(response, &telebot.SendOptions{
		ParseMode: "HTML",
	})
}

func (*StickerController) StickerCommand(ctx telebot.Context) error {
	args := ctx.Args()
	stickerID := args[0]

	response := &telebot.Sticker{
		File: telebot.File{
			FileID: stickerID,
		},
	}

	return ctx.Send(response)
}
