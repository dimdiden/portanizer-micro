package implementation

import (
	"context"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/dimdiden/portanizer-micro/services/users"

	"github.com/dimdiden/portanizer-micro/services/workbook"
	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// Service implements the Workbook service
type Service struct {
	users      users.Service
	repository workbook.Repository
	logger     log.Logger
}

// NewService creates and returns a new Workbook service instance
func NewService(users users.Service, repository workbook.Repository, logger log.Logger) workbook.Service {
	return &Service{
		users:      users,
		repository: repository,
		logger:     logger,
	}
}

func (s *Service) CreatePost(ctx context.Context, p workbook.Post) (*workbook.Post, error) {
	logger := log.With(s.logger, "method", "CreatePost")
	level.Debug(logger).Log("receivedPost", p.Title)

	claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
	// TODO: return correct error in case claims are not present
	if !ok {
		level.Error(logger).Log("err", "unexpected context", "contains", claims)
		return nil, workbook.ErrUnexpected
	}
	level.Debug(logger).Log("uid", claims.Subject)

	user, err := s.users.SearchByID(ctx, claims.Subject)
	if err != nil {
		return nil, users.ErrNotFound
	}
	p.UserID = user.ID

	post, err := s.repository.InsertPost(ctx, p)
	if err != nil {
		return nil, err
	}
	return post, nil
}
