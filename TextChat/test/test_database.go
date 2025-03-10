package main

import (
	"TextChat/src/database"
	"database/sql"
	"fmt"
)

func main() {
	var err error
	var inst *sql.DB
	// Test
	err = database.DBManager.Initialize("./db/TextChatDB.db")
	if err != nil {
		panic(err)
	}
	inst, err = database.DBManager.GetInstance()
	if err != nil {
		panic(err)
	}
	err = inst.Ping()
	if err != nil {
		panic(err)
	}
	// End
	fmt.Printf("Connection successfully established (status: %d).\n", inst.Stats().InUse)
}
