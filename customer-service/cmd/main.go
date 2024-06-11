package main

import (
	"fmt"
	"log"
	"net"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"customer-service/config"
	"customer-service/internal/handler"
	"customer-service/internal/repository"
	"customer-service/internal/usecase"
	"customer-service/lib/tracing"
	customer "customer-service/proto"
)

func main() {
	tracer, closer := tracing.InitTracing("customer-service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	cfg := config.NewConfig()
	config.Connect()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	customerRepo := repository.NewCustomerRepository(config.DB)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	s := grpc.NewServer()
	customer.RegisterCustomerServiceServer(s, customerHandler)

	fmt.Printf("Transaction Service is running on port %s\n", cfg.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
