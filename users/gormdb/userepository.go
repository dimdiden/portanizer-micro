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
		// users:  []*users.User{},
		logger: log.With(logger, "repository", "mock"),
	}
}

func (r *repository) InsertUser(ctx context.Context, email, pwd string) (*users.User, error) {
	// logger := log.With(r.logger, "method", "InsertUser")
	// for _, u := range r.users {
	// 	if u.Email != "" && u.Email == email {
	// 		level.Error(logger).Log("err", users.ErrExists)
	// 		return nil, users.ErrExists
	// 	}
	// }
	// fmt.Println("func (r *repository) InsertUser ", "ID=", string(len(r.users)+1))
	user := &users.User{ID: "2", Email: email, Password: pwd}
	r.users = append(r.users, user)
	return user, nil
	// return nil, users.ErrNotFound
}
