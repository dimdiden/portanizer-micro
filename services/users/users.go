package users

import "context"

type User struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password,omitempty" bson:"password"`
}

type Repository interface {
	InsertUser(ctx context.Context, email, pwd string) (*User, error)
	// SelectByCreds(ctx context.Context, email, pwd string) (*User, error)
}
