package dao

import "TextChat/src/model"

type ConversationDAO interface {
	CreateConversation(conversation *model.Conversation) (int, error)
	ReadConversation(conversationID int) (*model.Conversation, error)
	UpdateConversation(conversation *model.Conversation) error
	DeleteConversation(conversationID int) error
}
