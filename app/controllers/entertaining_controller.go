package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"gopkg.in/telebot.v3"
	"gorm.io/gorm"
	"n8body.dev/telebot-bot/app/dtos"
	"n8body.dev/telebot-bot/pkg/utils"
)

type EntertainingController struct {
	*gorm.DB
}

func (*EntertainingController) MemeCommand(ctx telebot.Context) error {
	args := ctx.Args()

	amount := 1
	if len(args) > 0 {
		if parsed, err := strconv.Atoi(args[0]); err == nil {
			amount = parsed
		}
	}

	url := fmt.Sprintf("https://meme-api.com/gimme/%d", amount)
	memeResponse, err := http.Get(url)
	if err != nil {
		return err
	}
	defer memeResponse.Body.Close()

	var memesData dtos.MemeResponseDTO
	if err := json.NewDecoder(memeResponse.Body).Decode(&memesData); err != nil {
		return err
	}

	for _, meme := range memesData.Memes {
		title := utils.EscapeHTML(meme.Title)
		source := fmt.Sprintf(`<tg-spoiler><a href="%s">source</a></tg-spoiler>`, meme.Preview[len(meme.Preview)-1])
		response := utils.JoinMultiline(title, source)

		if err := ctx.Send(response, &telebot.SendOptions{
			ParseMode: "HTML",
		}); err != nil {
			continue
		}
	}

	return nil
}
