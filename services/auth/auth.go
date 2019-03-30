package auth

import "context"

type Tokens struct {
	UID    string `json:"-"`
	Access string
}

type Repository interface {
	InsertTokens(ctx context.Context, UID string, tokens *Tokens) error
}
