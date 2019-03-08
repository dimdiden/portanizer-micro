package users

import "context"

type User struct {
	ID       string
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Repository interface {
	InsertUser(ctx context.Context, email, pwd string) (*User, error)
	// SelectByCreds(ctx context.Context, email, pwd string) (*User, error)
}
