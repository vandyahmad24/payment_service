package usecase

import (
	"api-gateway/internal/repository"
	"api-gateway/internal/request"
	"api-gateway/internal/response"
	"api-gateway/lib/tracing"
	"api-gateway/proto"
	"context"

	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (*response.CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, id string) (*proto.GetTransactionResponse, error)
	GetListTransactionByCustomerId(ctx context.Context, id string) (*proto.ListTransactionsResponse, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(repo repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: repo,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, req *request.CreateTransactionRequest) (*response.CreateTransactionResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	resp, err := u.transactionRepo.CreateTransaction(ctx, &proto.CreateTransactionRequest{
		Amount:        req.Amount,
		Currency:      req.Currency,
		PaymentMethod: req.PaymentMethod,
		Description:   req.Description,
		CustomerId:    req.UserId,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	response := &response.CreateTransactionResponse{
		TransactionId: resp.TransactionId,
		Status:        resp.Status,
		CreatedAt:     resp.CreatedAt,
	}

	return response, nil

}

func (u *transactionUsecase) GetTransaction(ctx context.Context, id string) (*proto.GetTransactionResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	resp, err := u.transactionRepo.GetTransaction(ctx, &proto.GetTransactionRequest{
		TransactionId: id,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	return resp, nil

}

func (u *transactionUsecase) GetListTransactionByCustomerId(ctx context.Context, id string) (*proto.ListTransactionsResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetListTransactionByCustomerId")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	resp, err := u.transactionRepo.ListTransactions(ctx, &proto.ListTransactionsRequest{
		CustomerId: id,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	if len(resp.Transactions) == 0 {
		return nil, status.Error(codes.NotFound, "Data not found")
	}

	return resp, nil

}
