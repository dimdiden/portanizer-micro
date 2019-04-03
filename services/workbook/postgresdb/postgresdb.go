package postgresdb

import (
	"context"
	"database/sql"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"

	"github.com/dimdiden/portanizer-micro/services/workbook"

	"github.com/lib/pq"
)

type Repository struct {
	db     *sql.DB
	logger log.Logger
}

// NewRepository returns a concrete repository
func NewRepository(connStr string, logger log.Logger) (workbook.Repository, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Repository{
		db:     db,
		logger: log.With(logger, "repository", "postgres"),
	}, nil
}

func (r *Repository) InsertPost(ctx context.Context, post workbook.Post) (*workbook.Post, error) {
	logger := log.With(r.logger, "method", "InsertPost")

	tx, err := r.db.Begin()
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	defer func() {
		if err != nil {
			level.Error(logger).Log("err", err)
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	var postID string
	err = tx.QueryRow("INSERT INTO posts (user_id, title, content) VALUES ($1, $2, $3) RETURNING post_id",
		post.UserID, post.Title, post.Content).Scan(&postID)
	level.Debug(logger).Log("post_id", postID)

	if err != nil {
		if e, ok := err.(*pq.Error); ok {
			switch e.Code.Class() {
			case "23":
				return nil, workbook.ErrExists
			default:
				level.Error(logger).Log("err", err)
				return nil, workbook.ErrQueryRepository
			}
		}
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}

	insertTagStmt, err := tx.Prepare("INSERT INTO tags (user_id, name) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT tag_is_unique_for_user DO NOTHING;")
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}

	for _, tag := range post.Tags {
		_, err = insertTagStmt.Exec(post.UserID, tag.Name)
		if err != nil {
			level.Error(logger).Log("err", err)
			return nil, workbook.ErrQueryRepository
		}
	}

	var tagarr []string
	for _, t := range post.Tags {
		tagarr = append(tagarr, t.Name)
	}

	selectTagStmt, err := tx.Prepare("SELECT tag_id, name FROM tags WHERE user_id = $1 AND name = ANY ($2)")
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	level.Debug(logger).Log("tagarr", strings.Join(tagarr[:], ","))

	rows, err := selectTagStmt.Query(post.UserID, pq.Array(tagarr))
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}

	var resulttag []workbook.Tag
	for rows.Next() {
		var t workbook.Tag
		if err := rows.Scan(&t.ID, &t.Name); err != nil {
			level.Error(logger).Log("err", err)
			return nil, workbook.ErrQueryRepository
		}
		level.Debug(logger).Log("tag", t.ID+";"+t.Name)
		resulttag = append(resulttag, t)
	}

	tagPostStmt, err := tx.Prepare("INSERT INTO posts_tags (post_id, tag_id) VALUES ($1, $2) ON CONFLICT ON CONSTRAINT post_tag_pkey DO NOTHING;")
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}

	for _, t := range resulttag {
		_, err := tagPostStmt.Exec(postID, t.ID)
		if err != nil {
			level.Error(logger).Log("err", err)
			return nil, workbook.ErrQueryRepository
		}
	}

	return &workbook.Post{ID: postID, Title: post.Title, Content: post.Content, Tags: resulttag}, nil
}
