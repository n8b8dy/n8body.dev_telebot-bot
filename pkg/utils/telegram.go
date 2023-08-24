package utils

import "gopkg.in/telebot.v3"

func IsChatOfRightType(chat *telebot.Chat, rightTypes ...telebot.ChatType) bool {
	for _, rightType := range rightTypes {
		if chat.Type == rightType {
			return true
		}
	}

	return false
}
