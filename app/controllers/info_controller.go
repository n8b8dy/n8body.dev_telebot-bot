package controllers

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"n8body.dev/telebot-bot/app/models"
	"n8body.dev/telebot-bot/pkg/utils"
)

type InfoController struct {
	*gorm.DB
}

func (controller *InfoController) StartCommand(ctx telebot.Context) error {
	sender := ctx.Sender()

	welcomeName := sender.FirstName
	if welcomeName == "" {
		welcomeName = sender.Username
	}

	response1 := fmt.Sprintf(
		"Welcome, %s! This is a bot made by n8body for educational and experimental purposes. The current version is written in Go. Use /help to see all the available commands.",
		welcomeName,
	)

	if err := ctx.Send(response1); err != nil {
		return err
	}

	var randSticker models.StartSticker
	controller.Raw("SELECT * FROM start_stickers ORDER BY RANDOM() LIMIT 1").Scan(&randSticker)

	response2 := &telebot.Sticker{
		File: telebot.File{
			FileID: randSticker.ID,
		},
	}

	return ctx.Send(response2)
}

func (*InfoController) HelpCommand(ctx telebot.Context) error {
	response := utils.JoinMultiline(
		"Here's a list of all the functionality of the bot:",
		"- Information Commands:",
		"  /start - shows initial message",
		"  /help, /commands - shows this message",
		"  /bot - shows information about the bot",
		"  /me - shows information about the user",
		"  /chat - shows information about the chat (cannot be used in private chats)",
		"- Sticker Commands:",
		"  [sticker] - returns the sticker's ID (can be used only in private chats)",
		"  /sticker [ID] - shows the sticker with the provided ID",
		"- Entertaining Commands:",
		"  /meme - sends a random meme",
		"  /photo - sends a random photo",
		"- Admin Commands:",
		"  (under construction)",
		"- Creator Commands... you don't need them",
		"- And some easter eggs",
	)

	return ctx.Send(response)
}

func (*InfoController) MeCommand(ctx telebot.Context) error {
	sender := ctx.Sender()

	firstName := sender.FirstName
	if firstName == "" {
		firstName = "_\\[DATA EXPUNGED\\]_"
	}

	lastName := sender.LastName
	if lastName == "" {
		lastName = "_\\[DATA EXPUNGED\\]_"
	}

	isBot := sender.IsBot
	// Happy Easter, Baen!
	if baenID, _ := strconv.ParseInt(os.Getenv("BAEN_ID"), 10, 64); sender.ID == baenID {
		isBot = true
	}

	responseCaption := fmt.Sprintf(
		utils.JoinMultiline(
			"Username: `%s`",
			"ID: `%d`",
			"First name: %s",
			"Last name: %s",
			"Bot: %t",
			"Premium: %t",
		),
		sender.Username,
		sender.ID,
		firstName,
		lastName,
		isBot,
		sender.IsPremium,
	)

	profilePhotos, err := ctx.Bot().ProfilePhotosOf(sender)
	if err != nil {
		return err
	} else if len(profilePhotos) == 0 {
		return ctx.Send(responseCaption, &telebot.SendOptions{
			ParseMode: "MarkdownV2",
		})
	}

	response := profilePhotos[0]
	response.Caption = responseCaption

	return ctx.Send(&response, &telebot.SendOptions{
		ParseMode: "MarkdownV2",
	})
}
