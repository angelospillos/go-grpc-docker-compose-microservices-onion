package main

import (
	"os"

	"github.com/joho/godotenv"

	"context"
	"log"

	pb "github.com/angelospillos/coinsorchestrator/proto"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	coinsOrchestratorHost := getEnv("COINS_ORCHESTRATOR_HOST", "localhost:7000")
	conn, err := grpc.Dial(coinsOrchestratorHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewCoinsOrchestratorClient(conn)

	response, err := client.GetTopCoinsByMarketCap(context.Background(), &pb.TopCoinsByMarketCapRequest{
		Limit: 10,
	})
	if err != nil {
		log.Fatalf("failed to fetch: %v", err)
	}

	log.Print(response)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
