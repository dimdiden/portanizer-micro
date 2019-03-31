package users

import "context"

type User struct {
	ID       string `bson:"_id,omitempty"`
	Email    string `bson:"email"`
	Password string `json:",omitempty" bson:"password"`
}

type Repository interface {
	InsertUser(ctx context.Context, email, pwd string) (*User, error)
	GetByCreds(ctx context.Context, email, pwd string) (*User, error)
	GetByID(ctx context.Context, id string) (*User, error)
}
