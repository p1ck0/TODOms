package repository

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/p1ck0/TODOms/pkg/models"
	mem "github.com/p1ck0/TODOms/pkg/repository/memdb"
)

type Repoistory interface {
	CreateTODO(ctx context.Context, todo models.TODO) error
	GetTODOs(ctx context.Context) ([]models.TODO, error)
	// SetTimeOutTODO(ctx context.Context, id string, timer time.Time) error
}

type Repo struct {
	TODO Repoistory
}

func NewRepo(db *memdb.MemDB) *Repo {
	return &Repo{
		TODO: mem.NewTODOrepo(db),
	}
}
