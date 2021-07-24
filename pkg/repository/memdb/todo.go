package memdb

import (
	"context"

	"github.com/hashicorp/go-memdb"
	"github.com/p1ck0/TODOms/pkg/models"
)

type TODOrepo struct {
	db *memdb.MemDB
}

func NewTODOrepo(db *memdb.MemDB) *TODOrepo {
	return &TODOrepo{
		db: db,
	}
}

func (r *TODOrepo) CreateTODO(ctx context.Context, todo models.TODO) error {
	txn := r.db.Txn(true)
	if err := txn.Insert("todo", todo); err != nil {
		return err
	}
	txn.Commit()
	return nil
}

func (r *TODOrepo) GetTODOs(ctx context.Context) ([]models.TODO, error) {
	txn := r.db.Txn(false)
	it, err := txn.Get("todo", "id")
	if err != nil {
		return nil, err
	}
	var todos []models.TODO

	for obj := it.Next(); obj != nil; obj = it.Next() {
		todos = append(todos, obj.(models.TODO))
	}

	return todos, nil

}
