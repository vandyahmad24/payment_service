package main

import (
	"fmt"
	"log"
	"net"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"transaction-service/config"
	"transaction-service/internal/handler"
	"transaction-service/internal/repository"
	"transaction-service/internal/usecase"
	"transaction-service/lib/constant"
	"transaction-service/lib/tracing"
	"transaction-service/proto"
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

	paymentMethodConn, err := grpc.NewClient(
		cfg.PAYMENT_METHOD_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer paymentMethodConn.Close()

	paymentClient := proto.NewPaymentMethodServiceClient(paymentMethodConn)

	transactionRepository := repository.NewTransactionRepository(config.DB, paymentClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	s := grpc.NewServer()
	proto.RegisterTransactionServiceServer(s, transactionHandler)

	fmt.Printf("Transaction Service is running on port %s\n", cfg.PORT)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
