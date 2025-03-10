package dao

import "TextChat/src/model"

type MessageDAO interface {
	CreateMessage(message *model.Message) error
	ReadMessage(messageID int) (*model.Message, error)
	UpdateMessage(message *model.Message) error
	DeleteMessage(messageID int) error
}
