package coins_test

import (
	"context"
	"testing"

	"google.golang.org/grpc"

	"github.com/stretchr/testify/assert"

	coins "github.com/angelospillos/coinsorchestrator"
	pricing "github.com/angelospillos/pricingservice/proto"
	ranking "github.com/angelospillos/rankingservice/proto"
)

type MockPricingService struct {
	PricingResponse pricing.PricingResponse
	Err             error
}

func (pr MockPricingService) GetPricing(ctx context.Context, in *pricing.PricingRequest, opts ...grpc.CallOption) (*pricing.PricingResponse, error) {
	return &pr.PricingResponse, pr.Err
}

type MockRankingService struct {
	RankingResponse ranking.RankingResponse
	Err             error
}

func (rk MockRankingService) GetRankings(ctx context.Context, in *ranking.RankingRequest, opts ...grpc.CallOption) (*ranking.RankingResponse, error) {
	return &rk.RankingResponse, rk.Err
}

func TestGetTopCoinsByMarketCap(t *testing.T) {

	mockPricingService := MockPricingService{pricing.PricingResponse{
		Quotes: map[string]float32{
			"BTC": 1.1,
			"ETH": 2.2,
			"BNB": 3.3,
			"DOGE": 4.4,
			"ADA": 5.5,
			"USDT": 6.6,
			"XRP": 7.7,
			"DOT": 8.8,
			"BCH": 9.9,
			"LTC": 10.10,
		},
	}, nil}

	mockRankingService := MockRankingService{ranking.RankingResponse{
		Ranks: []*ranking.Rank{
			{Order: 1, Symbol: "LTC"},
			{Order: 2, Symbol: "BCH"},
			{Order: 3, Symbol: "DOT"},
			{Order: 4, Symbol: "XRP"},
			{Order: 5, Symbol: "USDT"},
			{Order: 6, Symbol: "ADA"},
			{Order: 7, Symbol: "DOGE"},
			{Order: 8, Symbol: "BNB"},
			{Order: 9, Symbol: "ETH"},
			{Order: 10, Symbol: "BTC"},
		},
	}, nil}

	service := coins.Service{
		PricingService: mockPricingService,
		RankingService: mockRankingService,
	}

	result, err := service.GetTopCoinsByMarketCap(10)

	// With Testify
	assert.NoError(t, err)
	assert.Equal(t, "LTC", result[0].Symbol)
	assert.Equal(t, float32(10.10), result[0].Price)
	assert.Equal(t, int32(1), result[0].Rank)

	// Without Testify
	if result[0].Symbol!= "LTC" {
		t.Fatal("Symbol not what expected")
	}
	if result[0].Price!= 10.10 {
		t.Fatal("Price not what expected")
	}
	if result[0].Rank!= 1 {
		t.Fatal("Rank not what expected")
	}
}

func TestGetTopCoinsByMarketCapLimit(t *testing.T) {

	mockPricingService := MockPricingService{pricing.PricingResponse{
		Quotes: map[string]float32{},
	}, nil}

	mockRankingService := MockRankingService{ranking.RankingResponse{
		Ranks: []*ranking.Rank{},
	}, nil}

	service := coins.Service{
		PricingService: mockPricingService,
		RankingService: mockRankingService,
	}

	result, err := service.GetTopCoinsByMarketCap(1)

	assert.Empty(t, result)
	assert.Error(t, err)
}