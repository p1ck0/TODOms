package memdb

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/p1ck0/TODOms/pkg/models"
)

type TODOrepo struct {
	db *memdb.Txn
}

func NewTODOrepo(db *memdb.Txn) *TODOrepo {
	return &TODOrepo{
		db: db,
	}
}

func (r *TODOrepo) CreateTODO(ctx context.Context, todo models.TODO) error {
	if err := r.db.Insert("todo", todo); err != nil {
		return err
	}
	r.db.Commit()
	return nil
}

func (r *TODOrepo) GetTODOs(ctx context.Context) ([]models.TODO, error) {
	it, err := r.db.Get("todo", "id")
	if err != nil {
		return nil, err
	}
	var todos []models.TODO

	for obj := it.Next(); obj != nil; obj = it.Next() {
		todos = append(todos, obj.(models.TODO))
	}

	return todos, nil

}
