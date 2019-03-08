package workbook

import (
	"context"
	"errors"
)

var (
	// ErrNotFound is common error which occurs when an item is not found in the storage
	ErrNotFound = errors.New("item not found")
	// ErrExists is common error which occurs when there is a try
	// to save an item which already exists in the storage
	ErrExists = errors.New("item already exists")
	// ErrQueryRepository occurs when there is any other issue with the storage
	ErrQueryRepository = errors.New("unable to query repository")
)

// PostService is the workbook service to perform basic action on posts
type PostService interface {
	// Create(ctx context.Context, post Post) (string, error)
	// Update(ctx context.Context, pid string, post Post) error
	GetByID(ctx context.Context, pid string) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	// Delete(ctx context.Context, pid string) error
}

// TagService is the workbook service to perform basic action on tags
type TagService interface {
	// Create(ctx context.Context, pid string, tag Tag) (string, error)
	// Update(ctx context.Context, tid string, tag Tag) error
	// GetAll(ctx context.Context) ([]*Tag, error)
	// Delete(ctx context.Context, tid string) error
}
