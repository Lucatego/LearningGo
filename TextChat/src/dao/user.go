package dao

import (
	"TextChat/src/model"
)

type UserDAO interface {
	CreateUser(user *model.User) error
	ReadUser(userID int) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(userID int) error
}
