package implementation

import (
	"context"

	"github.com/dimdiden/portanizer_micro/services/workbook"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// postService implements the PostService
type postService struct {
	repository workbook.PostRepository
	logger     log.Logger
}

// NewPostService creates and returns a new Post service instance
func NewPostService(rep workbook.PostRepository, logger log.Logger) workbook.PostService {
	return &postService{
		repository: rep,
		logger:     logger,
	}
}

func (s *postService) Create(ctx context.Context, post workbook.Post) (string, error) {
	logger := log.With(s.logger, "method", "Create")

	id, err := s.repository.Create(ctx, post)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", workbook.ErrQueryRepository
	}
	return id, nil
}

func (s *postService) Update(ctx context.Context, pid string, post workbook.Post) error {
	logger := log.With(s.logger, "method", "Update")

	if err := s.repository.Update(ctx, pid, post); err != nil {
		level.Error(logger).Log("err", err)
		return workbook.ErrQueryRepository
	}
	return nil
}

func (s *postService) GetByID(ctx context.Context, pid string) (*workbook.Post, error) {
	logger := log.With(s.logger, "method", "GetByID")

	post, err := s.repository.GetByID(ctx, pid)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	return post, nil
}

func (s *postService) GetAll(ctx context.Context) ([]*workbook.Post, error) {
	logger := log.With(s.logger, "method", "GetAll")

	posts, err := s.repository.GetAll(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	return posts, nil
}

func (s *postService) Dalete(ctx context.Context, id string) error {
	logger := log.With(s.logger, "method", "Delete")

	if err := s.repository.Dalete(ctx, id); err != nil {
		level.Error(logger).Log("err", err)
		return workbook.ErrQueryRepository
	}
	return nil
}
