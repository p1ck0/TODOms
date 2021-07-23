package mem

import (
	"github.com/hashicorp/go-memdb"
)

func NewMemDB() (*memdb.Txn, error) {
	scheme := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": &memdb.TableSchema{
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	db, err := memdb.NewMemDB(scheme)
	if err != nil {
		return nil, err
	}
	return db.Txn(true), nil

}
