package memdb

import (
	"context"
	"time"

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

func (r *TODOrepo) SetTimeOutTODO(ctx context.Context, id string, timer time.Time) error {
	txn := r.db.Txn(false)
	raw, err := txn.First("todo", "id", id)
	if err != nil {
		return err
	}

	todo := raw.(models.TODO)
	todo.Timer.IsSet = true
	todo.Timer.IsTimeOut = false
	todo.Timer.Time = timer

	txn = r.db.Txn(true)

	if err := txn.Insert("todo", todo); err != nil {
		return err
	}
	txn.Commit()

	return nil
}

func (r *TODOrepo) OffTimeOutTODO(ctx context.Context, id string) error {
	txn := r.db.Txn(false)
	raw, err := txn.First("todo", "id", id)
	if err != nil {
		return err
	}

	todo := raw.(models.TODO)
	todo.Timer.IsSet = false
	todo.Timer.IsTimeOut = true

	txn = r.db.Txn(true)

	if err := txn.Insert("todo", todo); err != nil {
		return err
	}
	txn.Commit()

	return nil
}
