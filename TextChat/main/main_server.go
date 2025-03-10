package main

import (
	"TextChat/src/database"
	"TextChat/src/server"
)

const (
	ipAddress, port, protocol        = "127.0.0.1", "1080", "tcp"
	numberOfHandlers          uint64 = 4
	clientsPerHandler         uint64 = 4

	dbLocation = "./db/TextChatDB.db"
)

func main() {
	// Set Database
	err := database.DBManager.Initialize(dbLocation)
	if err != nil {
		panic(err)
	}
	// Open Server
	var s server.Server
	// Initialize
	s.SetServer(numberOfHandlers, clientsPerHandler)
	s.CreateSocket(ipAddress, port, protocol)
	// Run
	s.RunServer()
	// Close database
	err = database.DBManager.CloseDB()
	if err != nil {
		panic(err)
	}
}
