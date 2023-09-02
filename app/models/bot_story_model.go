package models

type BotStory struct {
	BaseModel
	Text           string
	UserTelegramID int64
	User           User `gorm:"foreignKey:UserTelegramID;references:TelegramID"`
}
