package coins

import (
	"context"
	"errors"

	pricing "github.com/angelospillos/pricingservice/proto"
	ranking "github.com/angelospillos/rankingservice/proto"
)

type Coin struct {
	Rank   int32
	Symbol string
	Price  float32
}

type Service struct {
	PricingService pricing.PricingServiceClient
	RankingService ranking.RankingServiceClient
}

func (svc Service) GetTopCoinsByMarketCap(limit int32) ([]Coin, error) {

	if limit < 10 {
		return nil, errors.New("limit not supported")
	}

	rankingRequest := ranking.RankingRequest{Limit: limit}
	rankingResponse, err := svc.RankingService.GetRankings(context.Background(), &rankingRequest)

	if err != nil {
		return nil, errors.New("ranking service unreachable" + err.Error())
	}

	var symbols []string
	for _, v := range rankingResponse.Ranks {
		symbols = append(symbols, v.Symbol)
	}

	pricingRequest := pricing.PricingRequest{Symbols: symbols}
	pricingResponse, err := svc.PricingService.GetPricing(context.Background(), &pricingRequest)

	if err != nil {
		return nil, errors.New("pricing service unreachable" + err.Error())
	}

	var toplist []Coin
	for _, v := range rankingResponse.Ranks {
		toplist = append(toplist, Coin{
			Rank:   v.Order,
			Symbol: v.Symbol,
			Price:  pricingResponse.Quotes[v.Symbol],
		})
	}

	return toplist, nil
}
