package implementation

import (
	"context"

	"github.com/dimdiden/portanizer_micro/services/workbook"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// tagService implements the TagService
type tagService struct {
	repository workbook.TagRepository
	logger     log.Logger
}

// NewTagService creates and returns a new Tag service instance
func NewTagService(rep workbook.TagRepository, logger log.Logger) workbook.TagService {
	return &tagService{
		repository: rep,
		logger:     logger,
	}
}

func (s *tagService) Create(ctx context.Context, pid string, tag workbook.Tag) (string, error) {
	logger := log.With(s.logger, "method", "Create")

	id, err := s.repository.Create(ctx, pid, tag)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", workbook.ErrQueryRepository
	}
	return id, nil
}

func (s *tagService) Update(ctx context.Context, tid string, tag workbook.Tag) error {
	logger := log.With(s.logger, "method", "Update")

	if err := s.repository.Update(ctx, tid, tag); err != nil {
		level.Error(logger).Log("err", err)
		return workbook.ErrQueryRepository
	}
	return nil
}

func (s *tagService) GetAll(ctx context.Context) ([]*workbook.Tag, error) {
	logger := log.With(s.logger, "method", "GetAll")

	tags, err := s.repository.GetAll(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	return tags, nil
}

func (s *tagService) Dalete(ctx context.Context, tid string) error {
	logger := log.With(s.logger, "method", "GetAll")

	if err := s.repository.Dalete(ctx, tid); err != nil {
		level.Error(logger).Log("err", err)
		return workbook.ErrQueryRepository
	}
	return nil
}
