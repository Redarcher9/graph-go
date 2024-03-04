package grpcHandler

import (
	pb "Gographql/proto"
	"context"
	"fmt"
	"log"
)

func (gh *GrpcHandler) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	_, _ = gh.mobileInteractor.GetMobilesByOs(ctx, "IOS")
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}

func (gh *GrpcHandler) GetMobilesByOs(ctx context.Context, req *pb.MobilesByOsRequest) (*pb.MobilesByOsResponse, error) {
	mobiles, err := gh.mobileInteractor.GetMobilesByOs(ctx, req.Os)
	if err != nil {
		log.Printf("%v", err)
	}

	var res []*pb.Mobile
	for _, mobile := range mobiles {
		res = append(res, &pb.Mobile{
			ModelID: mobile.ModelID,
			Name:    mobile.Name,
			Os:      mobile.OS,
			Country: mobile.Country,
			Brand: &pb.Mobile_Brand{
				Name:    mobile.Brand.Name,
				BrandID: mobile.Brand.BrandID,
			},
		})
	}
	return &pb.MobilesByOsResponse{
		Mobiles: res,
	}, nil
}

func (gh *GrpcHandler) mustEmbedUnimplementedGreetServiceServer() {
	fmt.Println("DO nothing")
}
