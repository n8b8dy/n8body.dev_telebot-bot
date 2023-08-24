package controllers

import "gorm.io/gorm"

type MainController struct {
	*InfoController
	*StickerController
	*EasterController
}

func NewMainController(db *gorm.DB) *MainController {
	return &MainController{
		InfoController:    &InfoController{db},
		StickerController: &StickerController{db},
		EasterController:  &EasterController{db},
	}
}
