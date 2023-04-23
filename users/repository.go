package users

import (
	"gin-practice/db"
)

type UserRepository interface {
	DropDatabase() error
	Create(user *User) error
	FindById(id int64) (User, error)
	Update(user *User) error
	DeleteById(id int64) error
}

type userRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) DropDatabase() error {
	err := r.db.QueryRow("DROP TABLE users;").Scan()
	return err
}

func (r *userRepository) Create(user *User) (err error) {
	stmt, err := r.db.Prepare(`
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

func (r *userRepository) FindById(id int64) (user User, err error) {
	user = User{}

	err = r.db.QueryRow(`
	SELECT 
	    id, first_name, last_name, email, password
	FROM
	    users
	WHERE
	    id=$1
	`, id).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	return
}

func (r *userRepository) Update(user *User) (err error) {
	stmt, err := r.db.Prepare(`
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

func (r *userRepository) DeleteById(id int64) (err error) {
	_, err = r.db.Exec(`
	DELETE FROM
	    users
	WHERE
	    id=$1
	`, id)

	return
}
