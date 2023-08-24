package models

type StartSticker struct {
	ID    string `gorm:"primaryKey"`
	Name  string
	Topic string
}
