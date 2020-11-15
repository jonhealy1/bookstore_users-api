package services

import (
	"github.com/jonhealy1/bookstore_/bookstore_users-api/domain/users"
	"github.com/jonhealy1/bookstore_/bookstore_users-api/utils/errors"
)

// CreateUser -
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, nil
}
