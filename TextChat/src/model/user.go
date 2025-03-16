package model

import "time"

type User struct {
	ID       int
	Username string
	Password string
	
	TimeCreated time.Time
}
