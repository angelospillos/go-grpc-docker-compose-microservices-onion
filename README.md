# go grpc docker compose microservices monorepo onion architecture protocol buffers
golang monorepo grpc proto buffers docker docker-compose microservices onion architecture

# Getting Started
These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See solution orchestration for notes on how to deploy the project on a live system.

Features include:

- API versioning allows you to alter behavior between different clients. ... Versioning is determined by the incoming client request, and may either be based on the request URL, or based on the request headers. There are a number of valid approaches to approaching versioning.
- Use of the Hexagonal Architecture to arrange the application into logical layers, with well-defined responsibilities.
- Golang-Based monorepo but also modular in case production deployment is needed

### Prerequisites

What things you need to install the software and how to install them


- [Visual Studio Code](https://code.visualstudio.com/)
- [Go](https://golang.org/dl/)

## Quick Start

```
git clone git@github.com:angelospillos/go-grpc-docker-compose-microservices-onion.git
cd go-grpc-docker-compose-microservices-onion
docker-compose build
docker-compose up
```
The server should now be running at `http://localhost:8080`.


# Concept Overview
The platform exposes an HTTP endpoint, which when fetched, displays an up-to-date list of top assets and their current prices in USD. 
* The endpoint supports `limit` parameter which indicates how many top coins should be returned.
* The output is `JSON`

Example call look like this:
```
$ curl http://127.0.0.1:8080/v1/top/coins/by-market-cap?limit=10
[
    {
        "Rank": 1,
        "Symbol": "BTC",
        "Price": 57438.984
    },
    {
        "Rank": 2,
        "Symbol": "ETH",
        "Price": 3885.3809
    },
    {
        "Rank": 3,
        "Symbol": "XRP",
        "Price": 1.5085038
    },
    {
        "Rank": 4,
        "Symbol": "BNB",
        "Price": 661.34
    },
    {
        "Rank": 5,
        "Symbol": "DOGE",
        "Price": 0.56885195
    },
    {
        "Rank": 6,
        "Symbol": "ADA",
        "Price": 1.7990011
    },
    {
        "Rank": 7,
        "Symbol": "USDT",
        "Price": 0.99994195
    },
    {
        "Rank": 8,
        "Symbol": "LINK",
        "Price": 50.532948
    },
    {
        "Rank": 9,
        "Symbol": "DOT",
        "Price": 40.0098
    },
    {
        "Rank": 10,
        "Symbol": "UNI",
        "Price": 38.970337
    }
]
```

## HTTP Status codes
Each response contains aproper HTTP codeset. Here are the details:
- 200 — Success
- 400 — Client error (e.g. validation)
- 500— Server error

## Errors

### Structure
In caseof an error (both client 4xx and server 5xx), response will contain an object with the following structure:
```
{
   "string"
}
```

The ranking and price information are always be up-to-date. For example let's say that Wings ranking changes from #9 to #20, the list should reflect that change.

### Data Sources

* Use [coinmarketcap API](https://coinmarketcap.com/api/) to get the current USD prices
* Use [cryptocompare API](https://www.cryptocompare.com/api#-api-data-coinlist-) to get the current ranking information.

### Software Architecture

The solution consist of 4 separate microservices that run independently (service oriented architecture):

* Cryptos Gateway - exposes a HTTP endpoints, example implemented that returns the up-to-date list of top coins prices.

* Coins Orchestrator - The bussiness layer regarding coins which uses any underline technical services (plugable) to fulfil bussines requirements.

* Pricing Service - keeps the up-to-date pricing information
* Ranking Service - keeps the up-to-date ranking information

## Microservice Structure

```
/src
    /cmd
    /proto
    /web
    example.go
```

The source folder contains sub-folders that arrange the application into logical
layers as suggested by the onion architecture.

-   `cmd` The cmd folder contains the commands needed to boostrap a server and an example client.

-   `web:` This is the adapter layer of the Hexagonal Architecture. It adapts
the HTTP transforms the HTTP requests from the external world to the service
layer and transforms the objects returned by the service layer to HTTP
responses.

-   `example.go`: The service layer coordinates high-level activities such as
creation of domain objects and asking them to perform tasks requested by the
external world. It interacts with the repository layer to save and restore objects.

-   `proto`: Protocol buffers are Google's language-neutral, platform-neutral, extensible mechanism for serializing structured data – think XML, but smaller, faster, and simpler. You define how you want your data to be structured once, then you can use special generated source code to easily write and read your structured data to and from a variety of data streams and using a variety of languages.

### Inter-Service Communication

#### GRPC
gRPC is a modern open source high performance Remote Procedure Call (RPC) framework that can run in any environment. It can efficiently connect services in and across data centers with pluggable support for load balancing, tracing, health checking and authentication. It is also applicable in last mile of distributed computing to connect devices, mobile applications and browsers to backend services.


#### REST API Versioned
A REST API (also known as RESTful API) is an application programming interface (API or web API) that conforms to the constraints of REST architectural style and allows for interaction with RESTful web services.

## Development


### Modular Development
Every microservice is seperated into layers to be able to support "plug and play" and easy Test Driven Development

Example tests are can be found in coinsorchestrator microservice.

### Generate a Proto

```
protoc --go_opt=paths=source_relative --go_out=plugins=grpc:. pricing.proto
```

### Add a new microservice in the monorepo

```
go mod edit -replace=example.com/microservicename=../microservicename
go mod tidy
```

###  Configuration
Every service loads an .env file that contains required configuration

## Solution Orchestration

Update the global env. file with the required API information.

```
docker-compose build
docker-compose up --remove-orphans
```

## Versioning

For the versions available, see the [tags on this repository](https://github.com/angelospillos/gateway/tags). 

## Authors

* **Angelos Pillos** - (https://www.angelospillos.com)

See also the list of [contributors](https://github.com/angelospillos/gateway/graphs/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* WATTx  https://wattx.io/