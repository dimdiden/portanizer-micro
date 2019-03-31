package users

import (
	"context"
	"errors"
)

var (
	// ErrNotFound is common error which occurs when a user is not found in the storage
	ErrNotFound = errors.New("user not found")
	// ErrPwd is incorrect password
	ErrPwd = errors.New("password is incorrect")
	// ErrExists is common error which occurs when there is a try
	// to save an item which already exists in the storage
	ErrExists = errors.New("user already exists")
	// ErrNotValid is the validation error
	ErrNotValid = errors.New("the value is not valid")
	// ErrUnexpected is general error which is not covered by unit test yet
	ErrUnexpected = errors.New("unexpected behavior")
	// ErrQueryRepository occurs when there is any other issue with the storage
	ErrQueryRepository = errors.New("unable to query users repository")
)

type Service interface {
	CreateAccount(ctx context.Context, email, pwd string) (*User, error)
	SearchByCreds(ctx context.Context, email, pwd string) (*User, error)
	SearchByID(ctx context.Context, id string) (*User, error)
}
