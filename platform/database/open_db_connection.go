package database

import (
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"n8body.dev/telebot-bot/app/models"
)

func OpenDBConnection() (*gorm.DB, error) {
	dsn := os.Getenv("DSN")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&models.StartSticker{}, &models.User{}, &models.BotStory{}); err != nil {
		return nil, err
	}

	return db, nil
}
