package models

type User struct {
	BaseModel
	TelegramID   int64 `gorm:"uniqueIndex"`
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsBot        bool
	IsPremium    bool
}
