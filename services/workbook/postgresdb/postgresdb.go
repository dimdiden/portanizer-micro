package postgresdb

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/services/workbook"
)

// type Repository struct {
// 	db     *sql.DB
// 	logger log.Logger
// }

// // NewRepository returns a concrete repository
// func NewRepository(connStr string, logger log.Logger) (workbook.Repository, error) {
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Repository{
// 		db:     db,
// 		logger: log.With(logger, "repository", "postgres"),
// 	}, nil
// }

// func (r *Repository) InsertPost(ctx context.Context, post workbook.Post) (*workbook.Post, error) {
// 	return nil, nil
// }

type Repository struct {
	logger log.Logger
}

// NewRepository returns a concrete repository
func NewRepository(connStr string, logger log.Logger) (workbook.Repository, error) {
	return &Repository{
		logger: log.With(logger, "repository", "postgres"),
	}, nil
}

func (r *Repository) InsertPost(ctx context.Context, post workbook.Post) (*workbook.Post, error) {
	return &workbook.Post{ID: "Fake", Title: post.Title, Content: post.Content, Tags: post.Tags}, nil
}
