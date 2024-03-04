package grpcHandler

import (
	"Gographql/internal/core/services"
	pb "Gographql/proto"
)

type GrpcHandler struct {
	mobileInteractor services.MobileInteractor
	pb.UnimplementedGreetServiceServer
}

func NewgrpcHandler(mobileInteractor services.MobileInteractor) *GrpcHandler {
	return &GrpcHandler{
		mobileInteractor: mobileInteractor,
	}
}
