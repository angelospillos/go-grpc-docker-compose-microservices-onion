package main

import (
	"os"

	"github.com/joho/godotenv"

	"log"
	"net"

	pricing "github.com/angelospillos/pricingservice"
	"github.com/angelospillos/pricingservice/cmcpro"
	pb "github.com/angelospillos/pricingservice/proto"
	"github.com/angelospillos/pricingservice/web"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pricingServicePort := getEnv("PRICING_SERVICE_PORT", ":6001")
	lis, err := net.Listen("tcp", pricingServicePort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cmcproApiUrl := getEnv("CMCPRO_API_URL", "https://pro-api.coinmarketcap.com")
	cmcproApiKey := getEnv("CMCPRO_API_KEY", "5bb0dda3-9935-4dc8-8eea-ee8df0160b52")
	cmcproapi := cmcpro.Api{URL: cmcproApiUrl, Key: cmcproApiKey}

	var grpcServer = grpc.NewServer()
	var webServer = web.Server{
		Service: pricing.Service{CMCProApi: cmcproapi},
	}
	pb.RegisterPricingServiceServer(grpcServer, webServer)
	grpcServer.Serve(lis)
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
