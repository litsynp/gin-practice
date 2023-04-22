package users

import (
	"gin-practice/src/db"
)

type IUserService interface {
	CreateUser(user User) (User, error)
}

type TUserService struct{}

var UserService *TUserService

func (s *TUserService) CreateUser(
	user User,
	createUser func(db *db.Database, user *User) error,
) (User, error) {
	err := createUser(db.GetDb(), &user)
	return user, err
}

func (s *TUserService) FindUserById(
	id int64,
	findUserById func(db *db.Database, id int64) (User, error),
) (User, error) {
	user, err := findUserById(db.GetDb(), id)
	return user, err
}

func (s *TUserService) UpdateUser(
	user User,
	updateUser func(db *db.Database, user *User) error,
) (User, error) {
	err := updateUser(db.GetDb(), &user)
	return user, err
}

func (s *TUserService) DeleteUserById(
	id int64,
	deleteUserById func(db *db.Database, id int64) error,
) error {
	err := deleteUserById(db.GetDb(), id)
	return err
}
