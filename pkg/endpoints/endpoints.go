package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/p1ck0/TODOms/pkg/service"
)

type Endpoints struct {
	Create     endpoint.Endpoint
	Get        endpoint.Endpoint
	SetTimeOut endpoint.Endpoint
}

func MakeEndpoints(s service.Serv) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		// SetTimeOut: makeSetTimeOut(s),
	}
}

func makeCreateEndpoint(s service.Serv) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.TODO.Create(ctx, req.TODO)
		return CreateResponse{ID: id, Err: err}, nil
	}
}

func makeGetEndpoint(s service.Serv) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetRequest)
		todos, err := s.TODO.Get(ctx)
		return GetResponse{TODOs: todos, Err: err}, nil
	}
}

// func makeSetTimeOut(s service.Service) endpoint.Endpoint {
// 	return func(ctx context.Context, request interface{}) (interface{}, error) {
// 		req := request.(SetTimeOutRequest)
// 		id, err := s.SetTimeOut(ctx, req.ID, req.Timer)
// 		return SetTimeResponse{ID: id, Err: err}, nil
// 	}
// }
