package repository

import (
	"context"

	"transaction-service/internal/models"
	"transaction-service/lib/tracing"
	"transaction-service/proto"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *models.Transaction) error
	FindById(ctx context.Context, id string) (*models.Transaction, error)
	GetByCustomerId(ctx context.Context, CustomerID string) ([]models.Transaction, error)
	FindPaymentMethod(ctx context.Context, req *proto.GetPaymentMethodRequest) (*proto.GetPaymentMethodResponse, error)
	FindPaymentMethodById(ctx context.Context, req *proto.GetPaymentMethodByIdRequest) (*proto.GetPaymentMethodResponse, error)
}

type transactionRepository struct {
	db            *gorm.DB
	clientPayment proto.PaymentMethodServiceClient
}

func NewTransactionRepository(db *gorm.DB, clientPayment proto.PaymentMethodServiceClient) TransactionRepository {
	return &transactionRepository{
		db:            db,
		clientPayment: clientPayment,
	}
}

func (r *transactionRepository) Create(ctx context.Context, tx *models.Transaction) error {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "TransactionRepository.Create")
	defer sp.Finish()

	tracing.LogRequest(sp, tx)

	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *transactionRepository) FindPaymentMethod(ctx context.Context, req *proto.GetPaymentMethodRequest) (*proto.GetPaymentMethodResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "FindPaymentMethod")
	defer span.Finish()

	tracing.LogRequest(span, req)

	resp, err := r.clientPayment.GetPaymentMethod(ctx, &proto.GetPaymentMethodRequest{
		MethodName: req.MethodName,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return resp, nil
}

func (r *transactionRepository) FindById(ctx context.Context, id string) (*models.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "FindById")
	defer span.Finish()

	tracing.LogRequest(span, id)

	var resp models.Transaction

	err := r.db.Where("id = ?", id).First(&resp).Error
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return &resp, nil

}

func (r *transactionRepository) FindPaymentMethodById(ctx context.Context, req *proto.GetPaymentMethodByIdRequest) (*proto.GetPaymentMethodResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "FindPaymentMethodById")
	defer span.Finish()

	tracing.LogRequest(span, req)

	resp, err := r.clientPayment.GetPaymentMethodById(ctx, &proto.GetPaymentMethodByIdRequest{
		Id: req.Id,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return resp, nil
}

func (r *transactionRepository) GetByCustomerId(ctx context.Context, CustomerID string) ([]models.Transaction, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "GetByCustomerId")
	defer span.Finish()

	tracing.LogRequest(span, CustomerID)

	var resp []models.Transaction

	err := r.db.Where("customer_id = ?", CustomerID).Find(&resp).Error
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return resp, nil
}
