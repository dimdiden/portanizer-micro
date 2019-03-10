package gormdb

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/users"
)

type repository struct {
	users  []*users.User
	logger log.Logger
}

func New(logger log.Logger) users.Repository {
	return &repository{
		logger: log.With(logger, "repository", "mock"),
	}
}

func (r *repository) InsertUser(ctx context.Context, email, pwd string) (*users.User, error) {
	user := &users.User{ID: "2", Email: email, Password: pwd}
	r.users = append(r.users, user)
	return user, nil
}
