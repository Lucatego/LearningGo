package model

import "time"

type Conversation struct {
	ID          int
	Title       string
	Description string

	TimeCreated time.Time
}
