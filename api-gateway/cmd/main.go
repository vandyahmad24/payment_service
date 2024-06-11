package main

import (
	"api-gateway/config"
	"api-gateway/internal/handler"
	"api-gateway/internal/repository"
	"api-gateway/internal/usecase"
	"api-gateway/lib/constant"
	middlewares "api-gateway/lib/middleware"
	"api-gateway/lib/tracing"
	"api-gateway/proto"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	r := gin.Default()

	tracer, closer := tracing.InitTracing(constant.TRACE_NAME)
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	cfg := config.NewConfig()

	transactionConn, err := grpc.NewClient(
		cfg.TRANSACTION_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer transactionConn.Close()

	customerConn, err := grpc.NewClient(
		cfg.CUSTOMER_SERVICE,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	defer customerConn.Close()

	customerClient := proto.NewCustomerServiceClient(customerConn)
	customerRepo := repository.NewCustomerRepository(customerClient)
	customerUsecase := usecase.NewCustomerUsecase(customerRepo)
	customerHandler := handler.NewCustomerHandler(customerUsecase)

	transactionClient := proto.NewTransactionServiceClient(transactionConn)
	transactionRepo := repository.NewTransactionRepository(transactionClient)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo)
	transactionHandler := handler.NewTransactionHandler(transactionUsecase)

	r.POST("/login", customerHandler.LoginCustomer)
	r.POST("/register", customerHandler.RegisterCustomer)
	r.POST("/transactions", middlewares.Auth(), transactionHandler.CreateTransaction)
	r.GET("/transactions/:id", middlewares.Auth(), transactionHandler.GetDetailTransaction)
	r.GET("/customers/:customerId/transactions", middlewares.Auth(), transactionHandler.GetListTransactionByCustomerId)

	r.Run(fmt.Sprintf(":%s", cfg.PORT))
}
