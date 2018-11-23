package workbook

import "context"

// Post represents an article or note or basically enithing that will have
// title, ID, some content and tags assigned to it
type Post struct {
	ID      string
	Title   string
	Content string
	Tags    []Tag
}

// Tag is a lable assigned to a post
type Tag struct {
	ID   string
	Name string
}

// PostRepository is the interaction with the post database
type PostRepository interface {
	Create(ctx context.Context, post Post) (string, error)
	Update(ctx context.Context, id string, post Post) error
	GetByID(ctx context.Context, id string) (*Post, error)
	GetAll(ctx context.Context) ([]*Post, error)
	Dalete(ctx context.Context, id string) error
}

// TagsRepository is the interaction with the tag database
type TagsRepository interface {
	Create(ctx context.Context, tag Tag) (string, error)
	Update(ctx context.Context, id string, tag Tag) error
	GetAll(ctx context.Context) ([]*Post, error)
	Dalete(ctx context.Context, id string) error
	// DaleteTags(ctx context.Context, ids []string) error
}
