package users

import (
	"gin-practice/src/db"
)

type IUserRepository interface {
	DropDatabase(db *db.Database) error
	Create(db *db.Database, user *User) error
	FindById(db *db.Database, id int64) (User, error)
	Update(db *db.Database, user *User) error
	DeleteById(db *db.Database, id int64) error
}

type TUserRepository struct{}

var UserRepository *TUserRepository

func (r *TUserRepository) DropDatabase(db *db.Database) error {
	err := db.QueryRow("DROP TABLE users;").Scan()
	return err
}

func (r *TUserRepository) Create(db *db.Database, user *User) (err error) {
	stmt, err := db.Prepare(`
	INSERT INTO users(
        first_name, last_name, email, password
	) 
	VALUES(
        $1, $2, $3, $4
	) RETURNING id
  	`)

	if err != nil {
		return
	}

	defer stmt.Close()
	err = stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.Password).Scan(&user.ID)

	return
}

func (r *TUserRepository) FindById(db *db.Database, id int64) (user User, err error) {
	user = User{}

	err = db.QueryRow(`
	SELECT 
	    id, first_name, last_name, email, password
	FROM
	    users
	WHERE
	    id=$1
	`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	return
}

func (r *TUserRepository) Update(db *db.Database, user *User) (err error) {
	stmt, err := db.Prepare(`
	UPDATE
	    users
	SET 
	    first_name=$1, last_name=$2, email=$3, password=$4
	WHERE
	    id=$5
	`)

	if err != nil {
		return
	}

	defer stmt.Close()
	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.ID)

	return
}

func (r *TUserRepository) DeleteById(db *db.Database, id int64) (err error) {
	_, err = db.Exec(`
	DELETE FROM
	    users
	WHERE
	    id=$1
	`, id)

	return
}
