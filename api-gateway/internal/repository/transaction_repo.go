package repository

import (
	"api-gateway/lib/tracing"
	transaction "api-gateway/proto"
	"context"

	"github.com/opentracing/opentracing-go"
)

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, req *transaction.CreateTransactionRequest) (*transaction.CreateTransactionResponse, error)
	GetTransaction(ctx context.Context, req *transaction.GetTransactionRequest) (*transaction.GetTransactionResponse, error)
	ListTransactions(ctx context.Context, req *transaction.ListTransactionsRequest) (*transaction.ListTransactionsResponse, error)
}

type transactionRepository struct {
	client transaction.TransactionServiceClient
}

func NewTransactionRepository(client transaction.TransactionServiceClient) TransactionRepository {
	return &transactionRepository{
		client: client,
	}
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, req *transaction.CreateTransactionRequest) (*transaction.CreateTransactionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransaction")
	defer span.Finish()

	tracing.LogRequest(span, req)

	transaction, err := r.client.CreateTransaction(ctx, &transaction.CreateTransactionRequest{
		Amount:        req.Amount,
		Currency:      req.Currency,
		PaymentMethod: req.PaymentMethod,
		Description:   req.Description,
		CustomerId:    req.CustomerId,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	tracing.LogResponse(span, transaction)

	return transaction, nil
}

func (r *transactionRepository) GetTransaction(ctx context.Context, req *transaction.GetTransactionRequest) (*transaction.GetTransactionResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "GetTransaction")
	defer span.Finish()

	tracing.LogRequest(span, req)

	transaction, err := r.client.GetTransaction(ctx, &transaction.GetTransactionRequest{
		TransactionId: req.TransactionId,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	tracing.LogResponse(span, transaction)

	return transaction, nil
}

func (r *transactionRepository) ListTransactions(ctx context.Context, req *transaction.ListTransactionsRequest) (*transaction.ListTransactionsResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "ListTransactions")
	defer span.Finish()

	tracing.LogRequest(span, req)

	transaction, err := r.client.ListTransactions(ctx, &transaction.ListTransactionsRequest{
		CustomerId: req.CustomerId,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	tracing.LogResponse(span, transaction)

	return transaction, nil
}
