package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"n8body.dev/telebot-bot/app/dtos"
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
		utils.JoinLine(
			`Welcome, %s! This is a bot made by n8body for educational and experimental purposes.`,
			`The current version is written in Go.`,
			`Use /help to see all the available commands.`,
		),
		welcomeName,
	)

	if err := ctx.Send(response1); err != nil {
		return err
	}

	var randSticker models.StartSticker
	controller.Raw("SELECT * FROM start_stickers ORDER BY RANDOM() LIMIT 1").Scan(&randSticker)

	response2 := &telebot.Sticker{
		File: telebot.File{
			FileID: randSticker.FileID,
		},
	}

	return ctx.Send(response2)
}

func (*InfoController) HelpCommand(ctx telebot.Context) error {
	response := utils.JoinMultiline(
		`Here's a list of all the functionality of the bot:`,
		`- Information Commands:`,
		`  /start - shows initial message`,
		`  /help, /commands - shows this message`,
		`  /bot - shows story about the bot (AI powered, xd)`,
		`  /me - shows information about the user`,
		`  /chat - shows information about the chat (cannot be used in private chats)`,
		`- Sticker Commands:`,
		`  [sticker] - returns the sticker's ID (can be used only in private chats)`,
		`  /sticker [ID] - shows the sticker with the provided ID`,
		`- Entertaining Commands:`,
		`  /meme [?:AMOUNT] - sends a random meme or the provided amount of memes`,
		`- Admin Commands:`,
		`  (under construction)`,
		`- Creator Commands... you don't need them`,
		`- And some easter eggs`,
	)

	return ctx.Send(response)
}

func (*InfoController) MeCommand(ctx telebot.Context) error {
	sender := ctx.Sender()

	firstName := utils.EscapeHTML(sender.FirstName)
	if firstName == "" {
		firstName = `<i>[DATA EXPUNGED]</i>`
	}

	lastName := utils.EscapeHTML(sender.LastName)
	if lastName == "" {
		lastName = `<i>[DATA EXPUNGED]</i>`
	}

	isBot := sender.IsBot
	// Happy Easter, Baen!
	if baenID, _ := strconv.ParseInt(os.Getenv("BAEN_ID"), 10, 64); sender.ID == baenID {
		isBot = true
	}

	responseCaption := fmt.Sprintf(
		utils.JoinMultiline(
			`Username: <code>%s</code>`,
			`ID: <code>%d</code>`,
			`First name: %s`,
			`Last name: %s`,
			`Bot: %t`,
			`Premium: %t`,
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
			ParseMode: "HTML",
		})
	}

	response := profilePhotos[0]
	response.Caption = responseCaption

	return ctx.Send(&response, &telebot.SendOptions{
		ParseMode: "HTML",
	})
}

func (controller *InfoController) BotCommand(ctx telebot.Context) error {
	sender := ctx.Message().Sender

	response1 := "Looking for the manuscripts in the archive... Wait a second..."
	if err := ctx.Send(response1); err != nil {
		return err
	}

	var botStory models.BotStory
	var response2 string

	if err := controller.Where(&models.BotStory{UserTelegramID: sender.ID}).First(&botStory).Error; err != nil {
		log.Println("Couldn't find a story for the sender in DB.")
	} else {
		response2 = botStory.Text

		return ctx.Send(response2, &telebot.SendOptions{
			ParseMode: "HTML",
		})
	}

	gptRequestBody, err := json.Marshal(dtos.GptRequestDTO{
		Model: "gpt-4",
		Messages: []dtos.GptMessageDTO{
			{"system", "You are a professional storyteller."},
			{"user", utils.JoinLine(
				"Write a brief story about a telegram bot named \"n8bodys_helper\".",
				"It is a personal bot of a developer who is known only as n8body. The bot is written in Golang.",
				"Select one of the following genres to tell the story: fairy tail, sci-fi, horror, crime, mystery, spy fiction, fantasy, western.",
				"Be creative about the plot. Words limit is around 300 words. Do not mention the genre.",
			)},
		},
		User: os.Getenv("OPENAI_USER_ID"),
	})
	if err != nil {
		return nil
	}

	log.Println("Asking ChatGPT for a story...")
	gptRequest, err := http.NewRequest(http.MethodPost, "https://api.openai.com/v1/chat/completions", bytes.NewReader(gptRequestBody))
	if err != nil {
		return err
	}
	defer gptRequest.Body.Close()

	gptRequest.Header.Set("Content-Type", "application/json")
	gptRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY")))

	gptResponse, err := http.DefaultClient.Do(gptRequest)
	if err != nil {
		return err
	}
	defer gptResponse.Body.Close()

	log.Println("Got the story!")

	var gptData dtos.GptResponseDTO
	if err := json.NewDecoder(gptResponse.Body).Decode(&gptData); err != nil {
		return err
	}

	botStory.User = models.User{
		TelegramID:   sender.ID,
		FirstName:    sender.FirstName,
		LastName:     sender.LastName,
		Username:     sender.Username,
		LanguageCode: sender.LanguageCode,
		IsBot:        sender.IsBot,
		IsPremium:    sender.IsPremium,
	}
	botStory.Text = gptData.Choices[0].Message.Content

	if err := controller.Create(&botStory).Error; err != nil {
		return err
	}

	response2 = botStory.Text

	return ctx.Send(response2, &telebot.SendOptions{
		ParseMode: "HTML",
	})
}
