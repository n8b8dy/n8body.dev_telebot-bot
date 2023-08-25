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
			FileID: randSticker.ID,
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
		`  /bot - shows information about the bot`,
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

func (*InfoController) BotCommand(ctx telebot.Context) error {
	// TODO: use OpenAI API to generate new story every time
	response := utils.JoinMultiline(
		`Once upon a time in the mystical realm of Cyberspace, there lived a clever and talented developer named n8body. n8body was renowned throughout the digital kingdom for their exceptional skills in the art of coding. They could craft intricate software spells, conjure elegant algorithms, and weave together lines of code with the grace of a wizard.`, ``,
		`One day, as n8body was deep in thought, pondering how to make their daily programming tasks easier, a magical idea struck them like a bolt of lightning. They decided to create a faithful companion, a Telegram bot named "n8bodys_helper." This bot would be their loyal assistant, ready to assist at any hour of the day or night.`, ``,
		`n8bodys_helper was no ordinary bot; it was crafted using the ancient and powerful programming language known as Golang. This language bestowed upon the bot a unique blend of speed, efficiency, and reliability, making it the ideal companion for a developer as skilled as n8body.`, ``,
		`This enchanted bot could perform an array of tasks that would make any developer's heart skip a beat. It could fetch the latest programming news, provide code snippets and explanations, generate random code challenges to sharpen n8body's skills, and even offer witty programming jokes to lighten the mood during long coding sessions.`, ``,
		`But n8bodys_helper's talents did not stop there. It was also a vigilant guardian of n8body's code, always on the lookout for bugs and errors. With a simple command, it would scan through lines of code, pointing out potential issues and suggesting improvements, much like a wise old sage guiding a young apprentice.`, ``,
		`n8body and n8bodys_helper quickly became inseparable companions. Together, they embarked on countless coding adventures, conquering complex projects and unraveling the mysteries of software development. With each passing day, n8body's skills grew stronger, and their reputation as a coding wizard soared higher.`, ``,
		`And so, in the mystical realm of Cyberspace, the tale of n8body and n8bodys_helper became the stuff of legend. It was a reminder that with a little creativity and the right tools, even the most challenging coding quests could be transformed into enchanting adventures.`, ``,
		`<tg-spoiler>the story was written by ChatGPT 3.5, xd</tg-spoiler>`,
	)

	return ctx.Send(response, &telebot.SendOptions{
		ParseMode: "HTML",
	})
}
