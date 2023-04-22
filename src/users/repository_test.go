package users

import (
	"gin-practice/src/db"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		defer setUpTest()()

		user := dummyUser()

		if err := UserRepository.Create(database, &user); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		var expectedId int64 = 1
		if user.ID != expectedId {
			t.Errorf("Expected user ID to be %v, got %v", expectedId, user.ID)
		}
	})

	t.Run("should return an error if email is not unique", func(t *testing.T) {
		defer setUpTest()()

		user := dummyUser()

		if err := UserRepository.Create(database, &user); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		user2 := dummyUser()
		err := UserRepository.Create(database, &user2)

		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

func TestFindUserById(t *testing.T) {
	t.Run("find user by ID", func(t *testing.T) {
		defer setUpTest()()

		existing := dummyUser()

		if err := UserRepository.Create(database, &existing); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		user, err := UserRepository.FindById(database, existing.ID)

		if err != nil {
			t.Errorf("Error finding user: %v", err)
		}

		if user.ID != existing.ID {
			t.Errorf("Expected user ID to be %v, got %v", existing.ID, user.ID)
		}
	})

	t.Run("should return an error if user does not exist", func(t *testing.T) {
		defer setUpTest()()

		_, err := UserRepository.FindById(database, 1)

		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update user by ID", func(t *testing.T) {
		defer setUpTest()()

		existing := dummyUser()

		if err := UserRepository.Create(database, &existing); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		var newEmail = "updated@example.com"
		updated := existing
		{
			updated.Email = newEmail
		}
		err := UserRepository.Update(database, &updated)

		if err != nil {
			t.Errorf("Error finding user: %v", err)
		}

		if updated.Email != newEmail {
			t.Errorf("Expected updated email to be %v, got %v", newEmail, updated.Email)
		}
	})
}

func TestDeleteById(t *testing.T) {
	t.Run("delete user by ID", func(t *testing.T) {
		defer setUpTest()()

		existing := dummyUser()

		if err := UserRepository.Create(database, &existing); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		err := UserRepository.DeleteById(database, existing.ID)

		if err != nil {
			t.Errorf("Error deleting user: %v", err)
		}

		_, err = UserRepository.FindById(database, existing.ID)

		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

var database *db.Database

func setUpTest() func() {
	database = db.GetDb()

	return func() {
		UserRepository.DropDatabase(database)
		db.CloseDb()
	}
}

func dummyUser() User {
	return User{
		Email:     "test@example.com",
		Password:  "password",
		FirstName: "John",
		LastName:  "Doe",
	}
}
