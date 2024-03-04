package grpcserver

import (
	"Gographql/internal/core/services"
	grpcHandler "Gographql/internal/handler/grpc"
	"Gographql/internal/infrastructure/repository"
	pb "Gographql/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8081"
)

func InitialiseGRPC() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to start the server")
	}
	grpcServer := grpc.NewServer()

	//Setup GRPC Server for Mobiles
	mobileRepository := repository.NewMobileRepo("Gorm")
	mobileService := services.NewMobileInteractor(&mobileRepository)
	mobileGrpcHandler := grpcHandler.NewgrpcHandler(*mobileService)
	pb.RegisterGreetServiceServer(grpcServer, mobileGrpcHandler)

	//Start GRPC server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start:%v", err)
	}
}
