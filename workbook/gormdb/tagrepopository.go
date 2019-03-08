package gormdb

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/jinzhu/gorm"
)

type tagRepository struct {
	db     *gorm.DB
	logger log.Logger
}

// NewTagRepository returns a concrete repository backed by gorm ORM
func NewTagRepository(db *gorm.DB, logger log.Logger) workbook.TagRepository {
	// return  repository
	return &tagRepository{
		db:     db,
		logger: log.With(logger, "repository", "gormdb"),
	}
}

func (r *tagRepository) Create(ctx context.Context, pid string, tag workbook.Tag) (string, error) {
	if !r.db.First(&tag, "name = ?", tag.Name).RecordNotFound() {
		return "", workbook.ErrExists
	}
	if err := r.db.Save(&tag).Error; err != nil {
		return "", err
	}
	return tag.ID, nil
}

func (r *tagRepository) Update(ctx context.Context, id string, tag workbook.Tag) error {
	if !r.db.First(&tag, "name = ?", tag.Name).RecordNotFound() && id != fmt.Sprint(tag.ID) {
		return workbook.ErrExists
	}
	var updTag workbook.Tag
	if r.db.First(&updTag, "id = ?", id).RecordNotFound() {
		return workbook.ErrNotFound
	}
	if err := r.db.Model(&updTag).Update(tag).Error; err != nil {
		return err
	}
	return nil
}

func (r *tagRepository) GetAll(ctx context.Context) ([]*workbook.Tag, error) {
	var tags []*workbook.Tag
	if err := r.db.Order("ID ASC").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *tagRepository) Dalete(ctx context.Context, id string) error {
	var tag workbook.Tag
	if r.db.First(&tag, "id = ?", id).RecordNotFound() {
		return workbook.ErrNotFound
	}
	if err := r.db.Delete(&tag).Error; err != nil {
		return err
	}
	return nil
}
