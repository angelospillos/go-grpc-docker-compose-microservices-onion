package ranking

import (
	"errors"

	"github.com/angelospillos/rankingservice/cryptocompare"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Rank struct {
	Order  int32
	Symbol string
}

type Service struct {
	CryptoCompareApi cryptocompare.Api
}

func (svc Service) GetRankings(limit int32) ([]Rank, error) {

	if limit < 10 {
		return nil, errors.New("not supported")
	}

	ranks, err := svc.CryptoCompareApi.GetTopByMarketCap(cryptocompare.TopByMarketCapRequest{
		Limit: limit,
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cryptocompare invalid argument %s", err.Error())
	}

	var response []Rank
	for k, v := range ranks.Data {
		response = append(response, Rank{
			Order:  int32(k) + 1, // we start from 0
			Symbol: v.CoinInfo.Name,
		})
	}

	return response, nil
}
