package models

import "github.com/google/uuid"

type BotStory struct {
	BaseModel
	Text   string
	UserID uuid.UUID
	User   User
}
