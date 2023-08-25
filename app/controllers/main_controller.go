package controllers

import "gorm.io/gorm"

type MainController struct {
	*InfoController
	*StickerController
	*EntertainingController
	*EasterController
}

func NewMainController(db *gorm.DB) *MainController {
	return &MainController{
		InfoController:         &InfoController{db},
		StickerController:      &StickerController{db},
		EntertainingController: &EntertainingController{db},
		EasterController:       &EasterController{db},
	}
}
