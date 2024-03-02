package grpcserver

import (
	pb "Gographql/proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

type helloServer struct {
	pb.GreetServiceServer
}

func InitialiseGRPC() {
	fmt.Println("GRPC server started!")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server")
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start:%v", err)
	}
	fmt.Println("GRPC server started!")
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
