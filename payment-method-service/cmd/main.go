package main

import (
	"fmt"
	"log"
	"net"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"

	"payment-method-service/config"
	"payment-method-service/internal/handler"
	"payment-method-service/internal/repository"
	"payment-method-service/internal/usecase"
	"payment-method-service/lib/constant"
	"payment-method-service/lib/tracing"
	transaction "payment-method-service/proto"
)

func main() {
	tracer, closer := tracing.InitTracing(constant.TRACE_NAME)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	cfg := config.NewConfig()
	config.Connect()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	transactionRepository := repository.NewPaymentMethodRepository(config.DB)
	transactionUsecase := usecase.NewPaymentMethodUsecase(transactionRepository)
	transactionHandler := handler.NewPaymentMethodHandler(transactionUsecase)

	s := grpc.NewServer()
	transaction.RegisterPaymentMethodServiceServer(s, transactionHandler)

	fmt.Printf("Transaction Service is running on port %s\n", cfg.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
