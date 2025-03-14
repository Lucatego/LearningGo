package service

import (
	"TextChat/src/model"
	"database/sql"
)

type ConversationSQLite struct {
	db *sql.DB
}

func (c ConversationSQLite) CreateConversation(conversation *model.Conversation) error {
	return nil
}

func (c ConversationSQLite) ReadConversation(conversationID int) (*model.Conversation, error) {
	var conversation = &model.Conversation{}

	return conversation, nil
}

func (c ConversationSQLite) UpdateConversation(conversation *model.Conversation) error {
	return nil
}

func (c ConversationSQLite) DeleteConversation(conversationID int) error {
	return nil
}
