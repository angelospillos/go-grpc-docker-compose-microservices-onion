package main

import (
	"os"

	"github.com/joho/godotenv"

	"context"
	"log"

	pb "github.com/angelospillos/pricingservice/proto"
	"google.golang.org/grpc"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pricingServiceHost := getEnv("PRICING_SERVICE_HOST", "localhost:6001")
	conn, err := grpc.Dial(pricingServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPricingServiceClient(conn)

	res, err := client.GetPricing(context.Background(), &pb.PricingRequest{
		Symbols: []string{"ETH", "BTC"},
	})
	if err != nil {
		log.Fatalf("failed to fetch: %v", err)
	}

	log.Print(res)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
