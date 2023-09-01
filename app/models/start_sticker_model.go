package models

type StartSticker struct {
	BaseModel
	FileID string
	Name   string
	Topic  string
}
