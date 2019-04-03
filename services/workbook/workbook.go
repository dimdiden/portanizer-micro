package workbook

import "context"

// Post represents an article or note or basically enithing that will have
// title, ID, some content and tags assigned to it
type Post struct {
	ID      string `json:"id"`
	UserID  string `json:"-"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags"`
}

// Tag is a lable assigned to a post
type Tag struct {
	ID     string `json:"id"`
	UserID string `json:"-"`
	Name   string `json:"name"`
}

// https://stackoverflow.com/questions/7296846/how-to-implement-one-to-one-one-to-many-and-many-to-many-relationships-while-de

type Repository interface {
	InsertPost(ctx context.Context, p Post) (*Post, error)
}
