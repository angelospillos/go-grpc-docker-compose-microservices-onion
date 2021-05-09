package main

import (
	"os"

	"github.com/joho/godotenv"

	"log"

	coinsorchestratorpb "github.com/angelospillos/coinsorchestrator/proto"
	cryptos "github.com/angelospillos/cryptosgateway"

	"github.com/angelospillos/cryptosgateway/web"
	"google.golang.org/grpc"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	coinsOrchestratorHost := getEnv("COINS_ORCHESTRATOR_HOST", "localhost:7000")
	coinsOrchestratorConn, err := grpc.Dial(coinsOrchestratorHost, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer coinsOrchestratorConn.Close()
	coinsOrchestrator := coinsorchestratorpb.NewCoinsOrchestratorClient(coinsOrchestratorConn)

	cryptosService := cryptos.Service{
		CoinsOrchestrator: coinsOrchestrator,
	}

	cryptosGatewayHost := getEnv("CRYPTOS_GATEWAY_HOST", ":8080")
	webServer := web.Server{CryptosService: cryptosService, Host: cryptosGatewayHost}
	webServer.Start()
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
