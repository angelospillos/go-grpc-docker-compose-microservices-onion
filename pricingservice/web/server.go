package web

import (
	"context"
	"github.com/angelospillos/pricingservice"
	pb "github.com/angelospillos/pricingservice/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedPricingServiceServer
	Service pricing.Service
}

func (s Server) GetPricing(ctx context.Context, req *pb.PricingRequest) (*pb.PricingResponse, error) {

	prices, err := s.Service.GetPrices(req.Symbols)

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument %s", err.Error())
	}

	res := &pb.PricingResponse{
		Quotes: prices,
	}

	return res, nil
}
