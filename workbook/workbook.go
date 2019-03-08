package workbook

import "context"

// Post represents an article or note or basically enithing that will have
// title, ID, some content and tags assigned to it
type Post struct {
	ID      string
	UserID  string `gorm:"unique;not null"`
	Title   string `gorm:"unique;not null"`
	Content string
	Tags    []Tag `gorm:"many2many:post_tags;"`
}

// Tag is a lable assigned to a post
type Tag struct {
	ID     string
	UserID string `gorm:"unique;not null"`
	Name   string `gorm:"unique;not null"`
}

// PostRepository is the interaction with the post database
type PostRepository interface {
	// Create(ctx context.Context, post Post) (string, error)
	// Update(ctx context.Context, id string, post Post) error
	GetByID(ctx context.Context, id string) (*Post, error)
	SelectAll(ctx context.Context) ([]*Post, error)
	// Dalete(ctx context.Context, id string) error
}

// TagRepository is the interaction with the tag database
type TagRepository interface {
	// Create(ctx context.Context, pid string, tag Tag) (string, error)
	// Update(ctx context.Context, id string, tag Tag) error
	// GetAll(ctx context.Context) ([]*Tag, error)
	// Dalete(ctx context.Context, id string) error
}
