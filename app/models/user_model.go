package models

type User struct {
	BaseModel
	TelegramID   int64
	FirstName    string
	LastName     string
	Username     string
	LanguageCode string
	IsBot        bool
	IsPremium    bool
}
