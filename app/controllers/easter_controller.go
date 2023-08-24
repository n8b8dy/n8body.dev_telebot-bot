package controllers

import (
	"math/rand"
	"regexp"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"n8body.dev/telebot-bot/pkg/utils"
)

type EasterController struct {
	*gorm.DB
}

func (*EasterController) DaPizda(ctx telebot.Context) error {
	chat := ctx.Chat()

	if !utils.IsChatOfRightType(chat, telebot.ChatGroup, telebot.ChatSuperGroup) {
		return nil
	}

	message := ctx.Message()

	match, err := regexp.MatchString(`(?mi)(?:^|\s)+[дd]+[аa]+[.!?:()]*$`, message.Text)
	if err != nil {
		return err
	} else if !match {
		return nil
	}

	// 10% chance
	if shouldReply := rand.Intn(10) > 8; !shouldReply {
		return nil
	}

	response := &telebot.Sticker{
		File: telebot.File{
			FileID: "CAACAgIAAxkBAAIJXGTkn-mf2u9UZQpX2csYQG5SUBZSAAJLAANcd6QTp0RihIYIK-gwBA",
		},
	}

	return ctx.Reply(response)
}
