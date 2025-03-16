package model

import "time"

type Message struct {
	ID       int
	Message  string
	TimeSend time.Time

	Sender       *User
	Conversation *Conversation
}
