package service

import (
	"TextChat/src/database"
	"TextChat/src/model"
	"database/sql"
	"time"
)

type UserSQLite struct {
	conn *sql.DB
}

func (u *UserSQLite) CreateUser(user *model.User) (int, error) {
	// Get a connection
	// TODO: Implement a connection pool
	if u.conn == nil {
		var err error
		u.conn, err = database.DBManager.GetInstance()
		if err != nil {
			return -1, err
		}
	}

	// Transaction
	tx, err := u.conn.Begin()
	if err != nil {
		return -1, err
	}
	defer tx.Rollback()

	// Create the statement
	stmt, err := tx.Prepare("INSERT INTO User (username, password) VALUES (?, ?);")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	// Create the user
	res, err := stmt.Exec(user.Username, user.Password)
	if err != nil {
		return -1, err
	}

	// Get the last inserted ID
	lastID, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}

	// Commit the changes
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return int(lastID), nil
}

func (u *UserSQLite) ReadUser(userID int) (*model.User, error) {
	// There is no necessity for a transaction here.
	var user = model.User{ID: userID}

	// Get a connection
	if u.conn == nil {
		var err error
		u.conn, err = database.DBManager.GetInstance()
		if err != nil {
			return nil, err
		}
	}

	// Query to select a row
	row := u.conn.QueryRow("SELECT username, time_created FROM User WHERE id = ?;", userID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	// Read the username
	var timeStr string
	err := row.Scan(&user.Username, &timeStr)
	if err != nil {
		return nil, err
	}

	// Parse the date
	user.TimeCreated, err = time.Parse("2006-01-02 15:04:05", timeStr)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserSQLite) UpdateUser(user *model.User) error {
	return nil
}

func (u *UserSQLite) DeleteUser(userID int) error {
	return nil
}
