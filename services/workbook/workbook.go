package workbook

import "context"

// Post represents an article or note or basically enithing that will have
// title, ID, some content and tags assigned to it
type Post struct {
	ID      string
	UserID  string `json:"-"`
	Title   string `gorm:"unique;not null"`
	Content string
	Tags    []Tag `gorm:"many2many:post_tags;"`
}

// Tag is a lable assigned to a post
type Tag struct {
	ID     string
	UserID string `json:"-"`
	Name   string `gorm:"unique;not null"`
}

// https://stackoverflow.com/questions/7296846/how-to-implement-one-to-one-one-to-many-and-many-to-many-relationships-while-de

type Repository interface {
	InsertPost(ctx context.Context, p Post) (*Post, error)
}
