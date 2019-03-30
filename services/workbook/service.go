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
	ErrQueryRepository = errors.New("unable to query workbook repository")
	// ErrUnexpected is general error which is not covered by unit test yet
	ErrUnexpected = errors.New("unexpected behavior")
)

type Service interface {
	CreatePost(ctx context.Context, p Post) (*Post, error)
}
