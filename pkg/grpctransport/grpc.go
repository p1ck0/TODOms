package grpctransport

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/p1ck0/TODOms/pkg/endpoints"
	"github.com/p1ck0/TODOms/pkg/models"
	"github.com/p1ck0/TODOms/pkg/pb"
)

type gRPCServer struct {
	create     gt.Handler
	get        gt.Handler
	setTimeOut gt.Handler
	pb.UnimplementedTODOServiceServer
}

func NewGRPCServer(endpoints endpoints.Endpoints, logger log.Logger) pb.TODOServiceServer {

	return &gRPCServer{
		create: gt.NewServer(
			endpoints.Create,
			decodeCreateTODORequest,
			encodeCreateTODOResponse,
		),
		get: gt.NewServer(
			endpoints.Get,
			decodeGetTODORequest,
			encodeGetTODOResponse,
		),
		setTimeOut: gt.NewServer(
			endpoints.SetTimeOut,
			decodeSetTimeOutTODORequest,
			encodeSetTimeOutTODOResponse,
		),
	}
}

func (s *gRPCServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	_, resp, err := s.create.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.CreateResponse), nil
}

func (s *gRPCServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	_, resp, err := s.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.GetResponse), nil
}

func (s *gRPCServer) SetTimeOut(ctx context.Context, req *pb.SetTimeOutRequest) (*pb.SetTimeResponse, error) {
	_, resp, err := s.setTimeOut.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.SetTimeResponse), nil
}

//func (*gRPCServer) mustEmbedUnimplementedTODOServiceServer() {}

func decodeCreateTODORequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateRequest)
	return endpoints.CreateRequest{TODO: models.TODO{
		Name: req.TODO.Name,
	}}, nil
}

func encodeCreateTODOResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.CreateResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.CreateResponse{ID: resp.ID}, nil
}

func decodeGetTODORequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetRequest)
	return req, nil
}

func encodeGetTODOResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GetResponse)
	var todos []*pb.TODO

	for _, todo := range resp.TODOs {
		todos = append(todos, &pb.TODO{
			ID:   todo.ID,
			Name: todo.Name,
			Timer: &pb.Timer{
				IsSet:     todo.Timer.IsSet,
				IsTimeOut: todo.Timer.IsTimeOut,
				Time:      todo.Timer.Time.Format(time.RFC3339),
			},
		})
	}
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.GetResponse{TODOs: todos}, nil
}

func decodeSetTimeOutTODORequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.SetTimeOutRequest)
	return endpoints.SetTimeOutRequest{ID: req.ID, Second: req.Second}, nil
}

func encodeSetTimeOutTODOResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.SetTimeResponse)
	if resp.Err != nil {
		return nil, resp.Err
	}
	return &pb.SetTimeResponse{ID: resp.ID}, nil
}
