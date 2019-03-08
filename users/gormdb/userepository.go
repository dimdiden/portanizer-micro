package gormdb

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/dimdiden/portanizer-micro/users"
)

type repository struct {
	users  []*users.User
	logger log.Logger
}

func New(logger log.Logger) users.Repository {
	return &repository{
		users:  make([]*users.User, 10),
		logger: log.With(logger, "repository", "mock"),
	}
}

func (r *repository) InsertUser(ctx context.Context, email, pwd string) (*users.User, error) {
	logger := log.With(r.logger, "method", "InsertUser")
	for _, u := range r.users {
		if u.Email == email {
			level.Error(logger).Log("err", users.ErrExists)
			return nil, users.ErrExists
		}
	}
	user := &users.User{ID: string(len(r.users) + 1)}
	r.users = append(r.users, user)
	return user, nil
}
