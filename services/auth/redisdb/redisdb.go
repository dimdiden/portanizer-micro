package redisdb

import (
	"context"
	"encoding/json"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/services/auth"
	"github.com/gomodule/redigo/redis"
)

type repository struct {
	conn   redis.Conn
	logger log.Logger
}

// TODO: need to Close conn somewhere
func NewRepository(address string, logger log.Logger) (*repository, error) {
	c, err := redis.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	return &repository{
		conn:   c,
		logger: log.With(logger, "repository", "redisdb"),
	}, nil
}

func (r *repository) InsertTokens(ctx context.Context, UID string, tokens *auth.Tokens) error {

	// serialize User object to JSON
	json, err := json.Marshal(tokens)
	if err != nil {
		return err
	}

	// SET object
	_, err = r.conn.Do("SET", UID, json)
	if err != nil {
		return err
	}

	return nil
}
