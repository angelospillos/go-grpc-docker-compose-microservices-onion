package main

import (
	"os"
	
	"context"
	"log"

	"github.com/joho/godotenv"

	pb "github.com/angelospillos/rankingservice/proto"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rankingServiceHost := getEnv("RANKING_SERVICE_HOST", "localhost:6000")
	conn, err := grpc.Dial(rankingServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRankingServiceClient(conn)

	response, err := client.GetRankings(context.Background(), &pb.RankingRequest{
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