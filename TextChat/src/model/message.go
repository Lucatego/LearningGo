package model

import "time"

type Message struct {
	id       int
	message  string
	timeSend time.Time

	sender *User
}
