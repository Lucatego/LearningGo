package main

import (
	"TextChat/src/dao"
	"TextChat/src/database"
	"TextChat/src/model"
	"TextChat/src/service"
	"fmt"
)

const (
	dbLocation = "./db/TextChatDB.db"
)

func main() {
	// Prepare
	err := database.DBManager.Initialize(dbLocation)
	if err != nil {
		panic(err)
	}

	// Create
	var userConn dao.UserDAO = &service.UserSQLite{}
	user := model.User{Username: "Test2", Password: "test2"}

	// Exec
	lastID, err := userConn.CreateUser(&user)
	if err != nil {
		panic(err)
	}
	user.ID = lastID

	// End
	fmt.Printf("User ID: %d\n", user.ID)
}
