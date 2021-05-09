package main

import (
	"os"

	"github.com/joho/godotenv"

	"log"
	"net"

	pricingpb "github.com/angelospillos/pricingservice/proto"
	rankingpb "github.com/angelospillos/rankingservice/proto"

	coins "github.com/angelospillos/coinsorchestrator"
	pb "github.com/angelospillos/coinsorchestrator/proto"
	"github.com/angelospillos/coinsorchestrator/web"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	rankingServiceHost := getEnv("RANKING_SERVICE_HOST", "localhost:6000")
	rankingServiceConn, err := grpc.Dial(rankingServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer rankingServiceConn.Close()
	rankingService := rankingpb.NewRankingServiceClient(rankingServiceConn)

	pricingServiceHost := getEnv("PRICING_SERVICE_HOST", "localhost:6001")
	pricingServiceConn, err := grpc.Dial(pricingServiceHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer pricingServiceConn.Close()
	pricingService := pricingpb.NewPricingServiceClient(pricingServiceConn)

	coinsService := coins.Service{
		PricingService: pricingService,
		RankingService: rankingService,
	}
	coinsOrchestratorPort := getEnv("COINS_ORCHESTRATOR_PORT", ":7000")
	lis, err := net.Listen("tcp", coinsOrchestratorPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var grpcServer = grpc.NewServer()
	var webServer = web.Server{Service: coinsService}
	pb.RegisterCoinsOrchestratorServer(grpcServer, webServer)
	grpcServer.Serve(lis)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
