package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"

	"github.com/p1ck0/TODOms/pkg/models"
	"github.com/p1ck0/TODOms/pkg/repository"
)

type TODOService struct {
	repository repository.Repo
	logger     log.Logger
}

func NewService(rep repository.Repo, logger log.Logger) *TODOService {
	return &TODOService{
		repository: rep,
		logger:     logger,
	}
}

func (s *TODOService) Create(ctx context.Context, todo models.TODO) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	todo.ID = uuid.New().String()
	if err := s.repository.TODO.CreateTODO(ctx, todo); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	return "ok", nil
}

func (s *TODOService) Get(ctx context.Context) ([]models.TODO, error) {
	logger := log.With(s.logger, "method", "Create")
	var (
		todos []models.TODO
		err   error
	)
	if todos, err = s.repository.TODO.GetTODOs(ctx); err != nil {
		level.Error(logger).Log("err", err)
		return []models.TODO{}, err
	}

	return todos, nil
}

// func (s *TODOService) SetTimeOut(ctx context.Context, id string, timer time.Time) (string, error) {
// 	logger := log.With(s.logger, "method", "Create")
// 	if err := s.repository.TODO.SetTimeOutTODO(ctx, id, timer); err != nil {
// 		level.Error(logger).Log("err", err)
// 		return "", err
// 	}

// 	return "ok", nil
// }
