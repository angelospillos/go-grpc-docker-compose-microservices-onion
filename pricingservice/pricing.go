package pricing

import (
	"errors"

	cmcpro "github.com/angelospillos/pricingservice/cmcpro"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	CMCProApi cmcpro.Api
}

func (s Service) GetPrices(symbols []string) (map[string]float32, error) {

	if len(symbols) < 1 {
		return nil, errors.New("empty list of symbols")
	}

	latestPrices, err := s.CMCProApi.GetLatestPricesBySymbols(cmcpro.GetLatestPricesBySymbolsRequest{
		Symbols: symbols,
	})

	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cmcpro invalid argument %s", err.Error())
	}

	response := make(map[string]float32)
	for _, v := range latestPrices.Data {
		response[v.Symbol] = v.Quote["USD"].Price
	}

	return response, nil
}
