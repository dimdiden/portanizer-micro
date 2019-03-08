package implementation

import (
	"context"

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
	logger := log.With(s.logger, "method", "GetAll")

	user, err := s.repository.InsertUser(ctx, email, pwd)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, users.ErrQueryRepository
	}
	return user, nil
}
