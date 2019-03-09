package implementation

import (
	"context"
	"fmt"

	"github.com/dimdiden/portanizer-micro/users"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	repository users.Repository
	logger     log.Logger
}

func NewService(repository users.Repository, logger log.Logger) users.Service {
	return &service{
		repository: repository,
		logger:     logger,
	}
}

func (s *service) CreateAccount(ctx context.Context, email, pwd string) (*users.User, error) {
	logger := log.With(s.logger, "method", "CreateAccount")

	user, err := s.repository.InsertUser(ctx, email, pwd)
	fmt.Println("func (s *service) CreateAccount(ctx context.Context, email, ", user.ID, user.Email, user.Password)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, users.ErrQueryRepository
	}
	return user, nil
}
