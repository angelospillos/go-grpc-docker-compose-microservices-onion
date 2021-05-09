package cryptos

import (
	"context"
	"errors"

	coinspb "github.com/angelospillos/coinsorchestrator/proto"
)

type Coin struct {
	Rank   int32
	Symbol string
	Price  float32
}

type Service struct {
	CoinsOrchestrator coinspb.CoinsOrchestratorClient
}

func (svc Service) GetTopCoinsByMarketCap(limit int32) ([]Coin, error) {

	if limit < 10 {
		return nil, errors.New("minimum limit is 10")
	}

	if limit > 100 {
		return nil, errors.New("maximum limit is 100")
	}

	topCoinsByMarketCapRequest := coinspb.TopCoinsByMarketCapRequest{Limit: limit}
	topCoinsByMarketCapResponse, err := svc.CoinsOrchestrator.GetTopCoinsByMarketCap(context.Background(), &topCoinsByMarketCapRequest)

	if err != nil {
		return nil, errors.New("coins orchestrator unreachable")
	}

	var response []Coin
	for _, v := range topCoinsByMarketCapResponse.Coins {
		response = append(response, Coin{
			Rank:   v.Rank,
			Symbol: v.Symbol,
			Price:  v.Price,
		})
	}

	return response, nil
}
