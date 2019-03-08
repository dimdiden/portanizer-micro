package gormdb

import (
	"context"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/jinzhu/gorm"
)

type postRepository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewPostRepository returns a concrete repository backed by gorm ORM
func NewPostRepository(db *gorm.DB, logger log.Logger) workbook.PostRepository {
	// return  repository
	return &postRepository{
		db:     db,
		logger: log.With(logger, "repository", "gormdb"),
	}
}

// func (r *postRepository) Create(ctx context.Context, post workbook.Post) (string, error) {
// 	if !r.db.First(&post, "title = ?", post.Title).RecordNotFound() {
// 		return "", workbook.ErrExists
// 	}

// 	newPost := workbook.Post{Title: post.Title, Content: post.Content}
// 	if err := r.db.Create(&newPost).Error; err != nil {
// 		return "", err
// 	}

// 	for _, t := range post.Tags {
// 		r.db.FirstOrCreate(&t, t)
// 		r.db.Model(&newPost).Association("Tags").Append(t)
// 	}
// 	return newPost.ID, nil
// }

// func (r *postRepository) Update(ctx context.Context, id string, post workbook.Post) error {
// 	if !r.db.First(&post, "title = ?", post.Title).RecordNotFound() && id != post.ID {
// 		return workbook.ErrExists
// 	}

// 	var updPost workbook.Post
// 	if r.db.First(&updPost, "id = ?", id).RecordNotFound() {
// 		return workbook.ErrNotFound
// 	}

// 	if err := r.db.Model(&updPost).Update(workbook.Post{Title: post.Title, Content: post.Content}).Error; err != nil {
// 		return err
// 	}
// 	// Create tag if doesn't exist and assign tags to post
// 	for _, t := range post.Tags {
// 		r.db.FirstOrCreate(&t, t)
// 		r.db.Model(&updPost).Association("Tags").Append(t)
// 	}
// 	return nil
// }

func (r *postRepository) GetByID(ctx context.Context, id string) (*workbook.Post, error) {
	// logger := log.With(r.logger, "method", "GetByID")

	var post workbook.Post
	if r.db.First(&post, "id = ?", id).RecordNotFound() {
		// level.Error(logger).Log("err", workbook.ErrNotFound)
		return nil, workbook.ErrNotFound
	}
	r.db.Preload("Tags").Order("ID ASC").Find(&post)

	return &post, nil
}

func (r *postRepository) SelectAll(ctx context.Context) ([]*workbook.Post, error) {
	var posts []*workbook.Post
	if err := r.db.Preload("Tags").Order("ID ASC").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

// func (r *postRepository) Dalete(ctx context.Context, id string) error {
// 	var post workbook.Post
// 	if r.db.First(&post, "id = ?", id).RecordNotFound() {
// 		return workbook.ErrNotFound
// 	}
// 	if err := r.db.Delete(&post).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
