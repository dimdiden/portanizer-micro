package implementation

import (
	"context"
	"regexp"

	jwt "github.com/dgrijalva/jwt-go"

	kitjwt "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

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
	logger := log.With(s.logger, "method", "SearchByID")

	if id == "" {
		claims, ok := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
		if !ok {
			level.Error(logger).Log("err", "unexpected context", "contains", claims)
			return nil, users.ErrUnexpected
		}
		level.Debug(logger).Log("uid", claims.Subject)
		id = claims.Subject
	}

	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
