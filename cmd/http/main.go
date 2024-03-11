package main

import (
	core "Gographql/cmd/core"
	grpcserver "Gographql/cmd/grpc"
	"Gographql/graph"
	generated "Gographql/graph/generated"
	"Gographql/internal/core/services"
	"Gographql/internal/handler/api"
	kafkaHandler "Gographql/internal/handler/kafka"
	"Gographql/internal/infrastructure/repository"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {

	//Set up Mobile Processor
	mobileRepository := repository.NewMobileRepo("Gorm")
	mobileService := services.NewMobileInteractor(&mobileRepository)
	mobileProcessor := api.NewMobileProcessor(mobileService)

	//Set up Brand Processor
	brandRepository := repository.NewBrandRepo("Gorm")
	brandService := services.NewBrandInteractor(&brandRepository)
	brandProcessor := api.NewBrandProcessor(brandService)

	//GraphQL Configuration
	graphConfig := generated.Config{Resolvers: &graph.Resolver{
		MobileProcessor: mobileProcessor,
		BrandProcessor:  brandProcessor,
	}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(graphConfig))

	//Instantiate Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", playground.Handler("GraphQL playground", "/query"))
	mux.HandleFunc("/query", srv.ServeHTTP)
	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	//Run GRPC server

	go func() {
		grpcserver.InitialiseGRPC()
	}()

	//Go Routine to run server
	go func() {
		fmt.Println("REST API started")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic(err)
		}
	}()

	//Go Routine to run Kafka
	//Initialise Processors for Kafka
	BrandHandlerKafka := kafkaHandler.NewBrandProcessor(brandService)

	go func() {
		core.InitKafkaProducer("mobiles", BrandHandlerKafka)
	}()

	// Prevent the main goroutine from exiting
	select {}
}
