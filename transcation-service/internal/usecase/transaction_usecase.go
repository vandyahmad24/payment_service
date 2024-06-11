package usecase

import (
	"context"
	"time"

	"transaction-service/internal/models"
	"transaction-service/internal/repository"
	"transaction-service/lib/tracing"
	"transaction-service/proto"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

type TransactionUsecase interface {
	CreateTransaction(ctx context.Context, req *proto.CreateTransactionRequest) (*models.Transaction, error)
	FindById(ctx context.Context, id string) (*models.TransactionResponse, error)
	GetByCustomerId(ctx context.Context, customerId string) (*proto.ListTransactionsResponse, error)
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
}

func NewTransactionUsecase(tr repository.TransactionRepository) TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: tr,
	}
}

func (u *transactionUsecase) CreateTransaction(ctx context.Context, req *proto.CreateTransactionRequest) (*models.Transaction, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "TransactionUsecase.CreateTransaction")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	//find payment method
	method, err := u.transactionRepo.FindPaymentMethod(ctx, &proto.GetPaymentMethodRequest{
		MethodName: req.PaymentMethod,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tx := &models.Transaction{
		Id:              uuid.New().String(),
		Amount:          req.Amount,
		Currency:        req.Currency,
		PaymentMethodId: method.Id,
		Description:     req.Description,
		CustomerID:      req.CustomerId,
		Status:          "pending",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	if err := u.transactionRepo.Create(ctx, tx); err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	return tx, nil
}

func (u *transactionUsecase) FindById(ctx context.Context, id string) (*models.TransactionResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "TransactionUsecase.FindById")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	//find payment method
	transaction, err := u.transactionRepo.FindById(ctx, id)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	method, err := u.transactionRepo.FindPaymentMethodById(ctx, &proto.GetPaymentMethodByIdRequest{
		Id: transaction.PaymentMethodId,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	resp := models.TransactionResponse{
		Id:            transaction.Id,
		Amount:        transaction.Amount,
		Currency:      transaction.Currency,
		PaymentMethod: method.MethodName,
		Description:   transaction.Description,
		CustomerID:    transaction.CustomerID,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
		UpdatedAt:     transaction.UpdatedAt,
	}

	return &resp, nil
}

func (u *transactionUsecase) GetByCustomerId(ctx context.Context, customerId string) (*proto.ListTransactionsResponse, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "TransactionUsecase.GetByCustomerId")
	defer sp.Finish()

	tracing.LogRequest(sp, customerId)

	resp, err := u.transactionRepo.GetByCustomerId(ctx, customerId)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	var respTemp []*proto.GetTransactionResponse

	for _, v := range resp {

		method, err := u.transactionRepo.FindPaymentMethodById(ctx, &proto.GetPaymentMethodByIdRequest{
			Id: v.PaymentMethodId,
		})
		if err != nil {
			tracing.LogError(sp, err)
			return nil, err
		}

		respTemp = append(respTemp, &proto.GetTransactionResponse{
			TransactionId: v.Id,
			Status:        v.Status,
			Amount:        v.Amount,
			Currency:      v.Currency,
			PaymentMethod: method.MethodName,
			Description:   v.Description,
			CustomerId:    v.CustomerID,
			UpdatedAt:     v.UpdatedAt.Format(time.RFC3339),
			CreatedAt:     v.CreatedAt.Format(time.RFC3339),
		})
	}

	return &proto.ListTransactionsResponse{Transactions: respTemp}, nil

}
