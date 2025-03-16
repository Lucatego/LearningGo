package service

import (
	"TextChat/src/database"
	"TextChat/src/model"
	"database/sql"
)

type UserSQLite struct {
	conn *sql.DB
}

func (u *UserSQLite) CreateUser(user *model.User) (int, error) {
	// Get a connection
	if u.conn == nil {
		var err error
		u.conn, err = database.DBManager.GetInstance()
		if err != nil {
			return -1, err
		}
	}

	// Create the user
	res, err := u.conn.Exec("INSERT INTO User (username, password) VALUES (?, ?);",
		user.Username, user.Password)
	if err != nil {
		return -1, err
	}

	// Get the last inserted ID
	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(lastID), nil
}

func (u *UserSQLite) ReadUser(userID int) (*model.User, error) {
	var result = &model.User{}

	return result, nil
}

func (u *UserSQLite) UpdateUser(user *model.User) error {
	return nil
}

func (u *UserSQLite) DeleteUser(userID int) error {
	return nil
}
