package users

import (
	"context"
	"errors"
)

var (
	// ErrExists is common error which occurs when there is a try
	// to save an item which already exists in the storage
	ErrExists = errors.New("item already exists")
	// ErrQueryRepository occurs when there is any other issue with the storage
	ErrQueryRepository = errors.New("unable to query repository")
)

type Service interface {
	CreateAccount(ctx context.Context, email, pwd string) (*User, error)
	// GetByCreds(ctx context.Context, email, pwd string) (*User, error)
}
