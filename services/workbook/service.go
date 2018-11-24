package workbook

import (
	"context"
	"errors"
)

var (
	ErrNotFound        = errors.New("item not found")
	ErrQueryRepository = errors.New("unable to query repository")
)

// PostService is the workbook service to perform basic action on posts
type PostService interface {
	Create(ctx context.Context, post Post) (string, error)
	Update(ctx context.Context, pid string, post Post) error
	GetByID(ctx context.Context, pid string) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	Dalete(ctx context.Context, pid string) error
}

// TagService is the workbook service to perform basic action on tags
type TagService interface {
	Create(ctx context.Context, pid string, tag Tag) (string, error)
	Update(ctx context.Context, tid string, tag Tag) error
	GetAll(ctx context.Context) ([]*Tag, error)
	Dalete(ctx context.Context, tid string) error
	// DaleteTags(ctx context.Context, ids []string) error
}
