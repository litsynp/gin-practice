package users

import "database/sql"

type IUserService interface {
	CreateUser(user User) (User, error)
}

type TUserService struct{}

var UserService *TUserService

func (s *TUserService) CreateUser(
	db *sql.DB,
	user User,
	createUser func(db *sql.DB, user *User) error,
) (User, error) {
	err := createUser(db, &user)
	return user, err
}

func (s *TUserService) FindUserById(
	db *sql.DB,
	id int64,
	findUserById func(db *sql.DB, id int64) (User, error),
) (User, error) {
	user, err := findUserById(db, id)
	return user, err
}

func (s *TUserService) UpdateUser(
	db *sql.DB,
	user User,
	updateUser func(db *sql.DB, user *User) error,
) (User, error) {
	err := updateUser(db, &user)
	return user, err
}

func (s *TUserService) DeleteUserById(
	db *sql.DB,
	id int64,
	deleteUserById func(db *sql.DB, id int64) error,
) error {
	err := deleteUserById(db, id)
	return err
}
