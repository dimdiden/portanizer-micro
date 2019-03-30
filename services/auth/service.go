package auth

import (
	"context"
	"errors"
)

var (
	// to save an item which already exists in the storage
	ErrIssueToken = errors.New("issue token service error")
)

type Service interface {
	IssueTokens(ctx context.Context, email, pwd string) (*Tokens, error)
}
