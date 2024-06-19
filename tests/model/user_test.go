package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go_api/internal/model"
)

func TestUserModel(t *testing.T) {
	// Test case: Create a new user
	user := model.User{
		Name:  "Alice",
		Email: "alice@example.com",
	}

	// Verify fields
	assert.Equal(t, "Alice", user.Name, "user name should be Alice")
	assert.Equal(t, "alice@example.com", user.Email, "user email should be alice@example.com")
}
