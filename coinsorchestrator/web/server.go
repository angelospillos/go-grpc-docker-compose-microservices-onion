package web

import (
	"context"

	"github.com/angelospillos/coinsorchestrator"
	pb "github.com/angelospillos/coinsorchestrator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedCoinsOrchestratorServer
	Service coins.Service
}

func (s Server) GetTopCoinsByMarketCap(ctx context.Context, req *pb.TopCoinsByMarketCapRequest) (*pb.TopCoinsByMarketCapResponse, error) {

	if req.Limit < 1 {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument %s", req.Limit)
	}

	coins, err := s.Service.GetTopCoinsByMarketCap(req.Limit)

	if err != nil {
		return nil, status.Errorf(codes.Aborted, err.Error())
	}

	var response pb.TopCoinsByMarketCapResponse
	for _, v := range coins {
		response.Coins = append(response.Coins, &pb.Coin{
			Rank:   v.Rank,
			Symbol: v.Symbol,
			Price:  v.Price,
		})
	}

	return &response, nil
}
