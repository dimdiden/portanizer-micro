package implementation

import (
	"github.com/dimdiden/portanizer_micro/services/auth"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/dgrijalva/jwt-go"
)

// service implements the Service
type service struct {
	users *pb.UserService
	repository auth.TokenPairRepository
	logger     log.Logger
}

// NewService creates and returns a new service instance
func NewService(rep auth.TokenPairRepository, logger log.Logger) auth.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

func (s *service) IssueTokens(ctx context.Context, email, pwd string) (*TokenPair, error) {
	logger := log.With(s.logger, "method", "IssueTokens")

	user, err := s.users.GetByCreds(email, pwd)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, auth.ErrNoUser
	}

	tp,err := issueTokens(user.ID)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, auth.ErrIssueToken
	}

	return tp, nil

}

func (s *service) issueTokens(uid string) (*TokenPair, error) {
	// Create the token
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * atokenExp).Unix(),
		Subject:   strconv.Itoa(int(user.ID)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	atoken, err := token.SignedString(h.asecret)
	if err != nil {
		return nil, err
	}
	// Change expires to generate refresh token
	claims.ExpiresAt = time.Now().Add(time.Hour * rtokenExp).Unix()
	rtoken, err := token.SignedString(h.rsecret)
	if err != nil {
		return nil, err
	}

	user.RToken = rtoken
	if err := h.userRepo.Refresh(user); err != nil {
		return nil, err
	}

	return &tokenPair{atoken, rtoken}, nil
}