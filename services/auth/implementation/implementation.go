package implementation

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dimdiden/portanizer-micro/services/auth"
	"github.com/dimdiden/portanizer-micro/services/users"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// service implements the Service
type service struct {
	us users.Service

	repository auth.Repository
	secret     string
	expire     time.Duration
	logger     log.Logger
}

// NewService creates and returns a new service instance
func NewService(us users.Service, repo auth.Repository, secret string, exp int, logger log.Logger) auth.Service {
	return &service{
		us:         us,
		repository: repo,
		secret:     secret,
		expire:     time.Duration(exp),
		logger:     logger,
	}
}

func (s *service) IssueTokens(ctx context.Context, email, pwd string) (*auth.Tokens, error) {
	logger := log.With(s.logger, "method", "IssueTokens")

	user, err := s.us.SearchByCreds(ctx, email, pwd)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	t, err := s.issueTokens(user.ID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, auth.ErrIssueToken
	}

	if err := s.repository.InsertTokens(ctx, user.ID, t); err != nil {
		level.Error(logger).Log("err", err)
		return nil, err
	}

	return t, nil

}

func (s *service) issueTokens(uid string) (*auth.Tokens, error) {
	// Create the token
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * s.expire).Unix(),
		Subject:   uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	Access, err := token.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}
	// Change expires to generate refresh token
	// claims.ExpiresAt = time.Now().Add(time.Hour * rtokenExp).Unix()
	// rtoken, err := token.SignedString(h.rsecret)
	// if err != nil {
	// 	return nil, err
	// }

	// user.RToken = rtoken
	// if err := h.userRepo.Refresh(user); err != nil {
	// 	return nil, err
	// }

	return &auth.Tokens{UID: uid, Access: Access}, nil
}
