module github.com/angelospillos/coinsorchestrator

go 1.16

require (
	github.com/angelospillos/pricingservice v0.0.0-00010101000000-000000000000
	github.com/angelospillos/rankingservice v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.3.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.37.0
	google.golang.org/protobuf v1.25.0
)

replace github.com/angelospillos/pricingservice => ../pricingservice

replace github.com/angelospillos/rankingservice => ../rankingservice
