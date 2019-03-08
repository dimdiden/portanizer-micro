package auth

import "context"

type Tokens struct {
	UserID string `gorm:"unique;not null" json:"uid"`
	AToken string `json:"atoken"`
}

type Repository interface {
	InsertTokens(ctx context.Context, userID string) (*Tokens, error)
}
