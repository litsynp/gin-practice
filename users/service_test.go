package users

import (
	"gin-practice/db"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Run("should create a new user", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		user, err := userService.CreateUser(dummyUser())
		if err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		var expectedId int64 = 1
		if user.ID != expectedId {
			t.Errorf("Expected user ID to be %v, got %v", expectedId, user.ID)
		}
	})

	t.Run("should return an error if email is not unique", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		_, err := userService.CreateUser(dummyUser())
		if err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		_, err = userService.CreateUser(dummyUser())
		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

func TestUserService_FindUserById(t *testing.T) {
	t.Run("find user by ID", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		existing, err := userService.CreateUser(dummyUser())
		if err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		user, err := userService.FindUserById(existing.ID)
		if err != nil {
			t.Errorf("Error finding user: %v", err)
		}

		if user.ID != existing.ID {
			t.Errorf("Expected user ID to be %v, got %v", existing.ID, user.ID)
		}
	})

	t.Run("should return an error if user does not exist", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		_, err := userService.FindUserById(1)
		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

func TestUserService_UpdateUser(t *testing.T) {
	t.Run("update user by ID", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		existing := dummyUser()

		if _, err := userService.CreateUser(existing); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		var newEmail = "updated@example.com"
		updated := existing
		{
			updated.Email = newEmail
		}
		updated, err := userService.UpdateUser(updated)

		if err != nil {
			t.Errorf("Error finding user: %v", err)
		}

		if updated.Email != newEmail {
			t.Errorf("Expected updated email to be %v, got %v", newEmail, updated.Email)
		}
	})
}

func TestUserService_DeleteUserById(t *testing.T) {
	t.Run("delete user by ID", func(t *testing.T) {
		defer setUpTest()()
		userService := NewUserService(NewUserRepository(db.GetDb()))

		existing := dummyUser()
		if _, err := userService.CreateUser(existing); err != nil {
			t.Errorf("Error creating user: %v", err)
		}

		err := userService.DeleteUserById(existing.ID)
		if err != nil {
			t.Errorf("Error deleting user: %v", err)
		}

		_, err = userService.FindUserById(existing.ID)
		if err == nil {
			t.Errorf("Expected error to be returned")
		}
	})
}

func setUpTest() func() {
	return func() {
		userRepository := NewUserRepository(db.GetDb())
		userRepository.DropDatabase()
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
