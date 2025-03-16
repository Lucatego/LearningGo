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
	fmt.Println("Testing an insertion")
	var name, password string
	fmt.Println("Username: ")
	fmt.Scanln(&name)
	fmt.Println("Password: ")
	fmt.Scanln(&password)

	var userConn dao.UserDAO = &service.UserSQLite{}
	user := model.User{Username: name, Password: password}

	// Exec
	lastID, err := userConn.CreateUser(&user)
	if err != nil {
		panic(err)
	}
	user.ID = lastID

	// End
	fmt.Printf("User ID: %d\n", user.ID)

	var searchID int
	fmt.Println("Search ID: ")
	fmt.Scanln(&searchID)
	res, err := userConn.ReadUser(searchID)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %d %s %s %s\n", res.ID, res.Username, res.Password, res.TimeCreated.String())

	// End
	err = database.DBManager.CloseDB()
	if err != nil {
		panic(err)
	}

	fmt.Println("Test successful")
}
