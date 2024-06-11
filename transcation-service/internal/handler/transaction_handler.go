package handler

import (
	"context"
	"time"

	"transaction-service/internal/usecase"
	"transaction-service/lib/tracing"
	transaction "transaction-service/proto"

	"github.com/opentracing/opentracing-go"
)

type TransactionHandler struct {
	usecase usecase.TransactionUsecase
	transaction.UnimplementedTransactionServiceServer
}

func NewTransactionHandler(u usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		usecase: u,
	}
}

func (h *TransactionHandler) CreateTransaction(ctx context.Context, req *transaction.CreateTransactionRequest) (*transaction.CreateTransactionResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.CreateTransaction(ctx, req)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &transaction.CreateTransactionResponse{
		TransactionId: tx.Id,
		Status:        tx.Status,
		CreatedAt:     tx.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (h *TransactionHandler) GetTransaction(ctx context.Context, req *transaction.GetTransactionRequest) (*transaction.GetTransactionResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.FindById(ctx, req.TransactionId)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &transaction.GetTransactionResponse{
		TransactionId: tx.Id,
		Status:        tx.Status,
		Amount:        int64(tx.Amount),
		Currency:      tx.Currency,
		PaymentMethod: tx.PaymentMethod,
		Description:   tx.Description,
		CustomerId:    tx.CustomerID,
		UpdatedAt:     tx.UpdatedAt.Format(time.RFC3339),
		CreatedAt:     tx.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (h *TransactionHandler) ListTransactions(ctx context.Context, req *transaction.ListTransactionsRequest) (*transaction.ListTransactionsResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.GetByCustomerId(ctx, req.CustomerId)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return tx, nil
}
