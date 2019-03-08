package implementation

import (
	"context"

	"github.com/dimdiden/portanizer-micro/workbook"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

// postService implements the PostService
type postService struct {
	repository workbook.PostRepository
	logger     log.Logger
}

// NewPostService creates and returns a new Post service instance
func NewPostService(repository workbook.PostRepository, logger log.Logger) workbook.PostService {
	return &postService{
		repository: repository,
		logger:     logger,
	}
}

// func (s *postService) Create(ctx context.Context, post workbook.Post) (string, error) {
// 	logger := log.With(s.logger, "method", "Create")

// 	id, err := s.repository.Create(ctx, post)
// 	if err != nil {
// 		level.Error(logger).Log("err", err)
// 		return "", workbook.ErrQueryRepository
// 	}
// 	return id, nil
// }

// func (s *postService) Update(ctx context.Context, pid string, post workbook.Post) error {
// 	logger := log.With(s.logger, "method", "Update")

// 	if err := s.repository.Update(ctx, pid, post); err != nil {
// 		level.Error(logger).Log("err", err)
// 		return workbook.ErrQueryRepository
// 	}
// 	return nil
// }

func (s *postService) GetByID(ctx context.Context, pid string) (*workbook.Post, error) {
	logger := log.With(s.logger, "method", "GetByID")

	// claims, _ := ctx.Value(kitjwt.JWTClaimsContextKey).(*jwt.StandardClaims)
	// fmt.Println("USER: ", claims.Subject)
	// fmt.Println(ctx)

	post, err := s.repository.GetByID(ctx, pid) // <= issue with error notification
	if err != nil {
		level.Error(logger).Log("err", err)
		switch {
		case err == workbook.ErrNotFound:
			return post, workbook.ErrNotFound
		default:
			return post, workbook.ErrQueryRepository
		}
	}
	return post, nil
}

func (s *postService) GetAll(ctx context.Context) ([]*workbook.Post, error) {
	logger := log.With(s.logger, "method", "GetAll")

	posts, err := s.repository.SelectAll(ctx)
	if err != nil {
		level.Error(logger).Log("err", err)
		return nil, workbook.ErrQueryRepository
	}
	return posts, nil
}

// func (s *postService) Delete(ctx context.Context, id string) error {
// 	logger := log.With(s.logger, "method", "Delete")

// 	if err := s.repository.Dalete(ctx, id); err != nil {
// 		level.Error(logger).Log("err", err)
// 		return workbook.ErrQueryRepository
// 	}
// 	return nil
// }

// tagService implements the TagService
type tagService struct {
	repository workbook.TagRepository
	logger     log.Logger
}

// NewTagService creates and returns a new Tag service instance
func NewTagService(repository workbook.TagRepository, logger log.Logger) workbook.TagService {
	return &tagService{
		repository: repository,
		logger:     logger,
	}
}

// func (s *tagService) Create(ctx context.Context, pid string, tag workbook.Tag) (string, error) {
// 	logger := log.With(s.logger, "method", "Create")

// 	id, err := s.repository.Create(ctx, pid, tag)
// 	if err != nil {
// 		level.Error(logger).Log("err", err)
// 		return "", workbook.ErrQueryRepository
// 	}
// 	return id, nil
// }

// func (s *tagService) Update(ctx context.Context, tid string, tag workbook.Tag) error {
// 	logger := log.With(s.logger, "method", "Update")

// 	if err := s.repository.Update(ctx, tid, tag); err != nil {
// 		level.Error(logger).Log("err", err)
// 		return workbook.ErrQueryRepository
// 	}
// 	return nil
// }

// func (s *tagService) GetAll(ctx context.Context) ([]*workbook.Tag, error) {
// 	logger := log.With(s.logger, "method", "GetAll")

// 	tags, err := s.repository.GetAll(ctx)
// 	if err != nil {
// 		level.Error(logger).Log("err", err)
// 		return nil, workbook.ErrQueryRepository
// 	}
// 	return tags, nil
// }

// func (s *tagService) Delete(ctx context.Context, tid string) error {
// 	logger := log.With(s.logger, "method", "GetAll")

// 	if err := s.repository.Dalete(ctx, tid); err != nil {
// 		level.Error(logger).Log("err", err)
// 		return workbook.ErrQueryRepository
// 	}
// 	return nil
// }
