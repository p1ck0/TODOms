package service

import (
	"context"
	"fmt"
	"time"

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
	todo.Timer.IsSet = false
	todo.Timer.IsTimeOut = false
	if err := s.repository.TODO.CreateTODO(ctx, todo); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	return todo.ID, nil
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

func (s *TODOService) SetTimeOut(ctx context.Context, id string, second uint64) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	timer := time.Now().Add(time.Second * time.Duration(second))
	if err := s.repository.TODO.SetTimeOutTODO(ctx, id, timer); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	go func() {
		s.setTimer(ctx, id, timer)
	}()

	return "ok", nil
}

type TimeExp struct {
	ID    string
	Timer *time.Timer
}

func (s *TODOService) setTimer(ctx context.Context, id string, t time.Time) {
	fmt.Println("timer start")
	exp := t.Sub(time.Now())
	timeExp := TimeExp{
		ID:    id,
		Timer: time.NewTimer(exp),
	}
	<-timeExp.Timer.C
	fmt.Println("timer stop")
	if err := s.repository.TODO.OffTimeOutTODO(ctx, id); err != nil {
		fmt.Println(err)
	}
}
