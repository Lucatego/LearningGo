package service

import (
	"TextChat/src/model"
	"database/sql"
)

type UserSQLite struct {
	db *sql.DB
}

func (u UserSQLite) CreateUser(user *model.User) error {
	return nil
}

func (u UserSQLite) ReadUser(userID int) (*model.User, error) {
	var result = &model.User{}

	return result, nil
}

func (u UserSQLite) UpdateUser(user *model.User) error {
	return nil
}

func (u UserSQLite) DeleteUser(userID int) error {
	return nil
}
