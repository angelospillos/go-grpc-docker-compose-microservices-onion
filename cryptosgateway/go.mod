module github.com/angelospillos/cryptosgateway

go 1.16

require (
	github.com/angelospillos/coinsorchestrator v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	google.golang.org/grpc v1.37.0
)

replace github.com/angelospillos/coinsorchestrator => ../coinsorchestrator

replace github.com/angelospillos/pricingservice => ../pricingservice

replace github.com/angelospillos/rankingservice => ../rankingservice
