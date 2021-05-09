package main

import (
	"os"

	"github.com/joho/godotenv"

	"log"
	"net"

	ranking "github.com/angelospillos/rankingservice"
	"github.com/angelospillos/rankingservice/cryptocompare"
	pb "github.com/angelospillos/rankingservice/proto"
	"github.com/angelospillos/rankingservice/web"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rankingServicePort := getEnv("RANKING_SERVICE_PORT", ":6000")
	lis, err := net.Listen("tcp", rankingServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cryptocompareApiUrl := getEnv("CRYPTOCOMPARE_API_URL", "https://min-api.cryptocompare.com")
	cryptocompareApiKey := getEnv("CRYPTOCOMPARE_API_KEY", "4d9e2793fef2d55c81f68edc46a6f192378eec1184edc275cccfcc34eb89743e")
	cryptocompareApiBaseCurrency := getEnv("CRYPTOCOMPARE_API_BASE_CURRENCY", "USD")
	cryptocompareApi := cryptocompare.Api{URL: cryptocompareApiUrl, Key: cryptocompareApiKey, BaseCurrency: cryptocompareApiBaseCurrency}

	var grpcServer = grpc.NewServer()
	var webServer = web.Server{
		Service: ranking.Service{
			CryptoCompareApi: cryptocompareApi,
		},
	}
	pb.RegisterRankingServiceServer(grpcServer, webServer)
	grpcServer.Serve(lis)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
