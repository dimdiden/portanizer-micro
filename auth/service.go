package auth

import (
	"context"
)

type Service interface {
	IssueTokens(ctx context.Context, email, pwd string) (*Tokens, error)
}
