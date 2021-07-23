package service

import (
	"context"
	"errors"

	"github.com/go-kit/kit/log"

	"github.com/p1ck0/TODOms/pkg/models"
	"github.com/p1ck0/TODOms/pkg/repository"
)

type Service interface {
	Create(ctx context.Context, todo models.TODO) (string, error)
	Get(ctx context.Context) ([]models.TODO, error)
	// SetTimeOut(ctx context.Context, id string, timer time.Time) (string, error)
}

type Serv struct {
	TODO Service
}

func NewServ(r repository.Repo, logger log.Logger) *Serv {
	return &Serv{
		TODO: NewService(r, logger),
	}
}

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)
