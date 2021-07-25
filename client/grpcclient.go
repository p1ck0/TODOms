package main

import (
	"context"
	"fmt"

	"github.com/p1ck0/TODOms/pkg/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(":8081", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	client := pb.NewTODOServiceClient(conn)
	requestcreate := &pb.CreateRequest{
		TODO: &pb.TODO{
			Name: "robert",
		},
	}
	//requestget := &pb.GetRequest{}
	// requestset := &pb.SetTimeOutRequest{
	// 	ID:     "1a32bd50-c0f3-4c39-b5e1-63b17a21f963",
	// 	Second: 30,
	// }
	response, err := client.Create(context.Background(), requestcreate)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	fmt.Println(response)
}
