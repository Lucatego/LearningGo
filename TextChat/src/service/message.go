package service

import (
	"TextChat/src/model"
	"database/sql"
)

type MessageSQLite struct {
	db *sql.DB
}

func (u MessageSQLite) CreateMessage(message *model.Message) error {
	return nil
}

func (u MessageSQLite) ReadMessage(messageID int) (*model.Message, error) {
	var result = &model.Message{}

	return result, nil
}

func (u MessageSQLite) UpdateMessage(message *model.Message) error {
	return nil
}

func (u MessageSQLite) DeleteMessage(messageID int) error {
	return nil
}
