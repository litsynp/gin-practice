package users

import (
	"database/sql"
)

type IUserRepository interface {
	Migrate(db *sql.DB) error
	DropDatabase(db *sql.DB) error
	Create(db *sql.DB, user *User) error
	FindById(db *sql.DB, id int64) (User, error)
	Update(db *sql.DB, user *User) error
	DeleteById(db *sql.DB, id int64) error
}

type TUserRepository struct{}

var UserRepository *TUserRepository

func (r *TUserRepository) Migrate(db *sql.DB) error {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	    id SERIAL PRIMARY KEY,
	    first_name VARCHAR(20),
	    last_name VARCHAR(20),
	    email TEXT NOT NULL UNIQUE,
	    password VARCHAR(255)
 	);
	`)

	return err
}

func (r *TUserRepository) DropDatabase(db *sql.DB) error {
	err := db.QueryRow("DROP TABLE users;").Scan()
	return err
}

func (r *TUserRepository) Create(db *sql.DB, user *User) (err error) {
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

func (r *TUserRepository) FindById(db *sql.DB, id int64) (user User, err error) {
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

func (r *TUserRepository) Update(db *sql.DB, user *User) (err error) {
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

func (r *TUserRepository) DeleteById(db *sql.DB, id int64) (err error) {
	_, err = db.Exec(`
	DELETE FROM
	    users
	WHERE
	    id=$1
	`, id)

	return
}
