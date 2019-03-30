package implementation

import (
	"context"
	"regexp"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/services/users"
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
	// logger := log.With(s.logger, "method", "CreateAccount")

	if !isEmailValid(email) {
		return nil, users.ErrNotValid
	}
	if len(pwd) < 9 {
		return nil, users.ErrNotValid
	}

	user, err := s.repository.InsertUser(ctx, email, pwd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// based on http://www.golangprograms.com/regular-expression-to-validate-email-address.html
func isEmailValid(email string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

func (s *service) SearchByCreds(ctx context.Context, email, pwd string) (*users.User, error) {
	// logger := log.With(s.logger, "method", "SearchByCreds")

	if !isEmailValid(email) {
		return nil, users.ErrNotValid
	}

	user, err := s.repository.GetByCreds(ctx, email, pwd)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *service) SearchByID(ctx context.Context, id string) (*users.User, error) {
	// logger := log.With(s.logger, "method", "SearchByID")

	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
