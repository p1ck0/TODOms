package endpoints

import (
	"github.com/p1ck0/TODOms/pkg/models"
)

type CreateRequest struct {
	TODO models.TODO
}

type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

type GetRequest interface{}

type GetResponse struct {
	TODOs []models.TODO `json:"todos"`
	Err   error         `json:"error,omitempty"`
}

type SetTimeOutRequest struct {
	ID     string `json:"id"`
	Second uint64 `json:"timer"`
}

type SetTimeResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}
