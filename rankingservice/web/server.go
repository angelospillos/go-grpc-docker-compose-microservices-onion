package web

import (
	"context"

	"github.com/angelospillos/rankingservice"
	pb "github.com/angelospillos/rankingservice/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedRankingServiceServer
	Service ranking.Service
}

func (s Server) GetRankings(ctx context.Context, req *pb.RankingRequest) (*pb.RankingResponse, error) {

	rankings, err := s.Service.GetRankings(req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Argument %s", req.Limit)
	}

	var response pb.RankingResponse
	for _, v := range rankings {
		response.Ranks = append(response.Ranks, &pb.Rank{
			Order:  v.Order,
			Symbol: v.Symbol,
		})
	}

	return &response, nil
}
