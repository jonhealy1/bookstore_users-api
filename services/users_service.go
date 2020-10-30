package services

import (
	"github.com/jonhealy1/bookstore_/bookstore_users-api/domain/users"
)

// CreateUser -
func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
