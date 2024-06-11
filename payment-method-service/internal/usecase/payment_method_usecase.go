package usecase

import (
	"context"

	"payment-method-service/internal/models"
	"payment-method-service/internal/repository"
	"payment-method-service/lib/tracing"
	"payment-method-service/proto"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
)

type PaymentMethodUsecase interface {
	CreatePaymentMethod(ctx context.Context, req *proto.CreatePaymentMethodRequest) (*models.PaymentMethod, error)
	GetPaymentMethod(ctx context.Context, methodName string) (*models.PaymentMethod, error)
	GetPaymentMethodById(ctx context.Context, id string) (*models.PaymentMethod, error)
}

type paymentMethodUsecase struct {
	transactionRepo repository.PaymentMethodRepository
}

func NewPaymentMethodUsecase(tr repository.PaymentMethodRepository) PaymentMethodUsecase {
	return &paymentMethodUsecase{
		transactionRepo: tr,
	}
}

func (u *paymentMethodUsecase) CreatePaymentMethod(ctx context.Context, req *proto.CreatePaymentMethodRequest) (*models.PaymentMethod, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodUsecase.CreatePaymentMethod")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx := &models.PaymentMethod{
		Id:         uuid.New().String(),
		MethodName: req.MethodName,
	}

	if err := u.transactionRepo.Create(ctx, tx); err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	return tx, nil
}

func (u *paymentMethodUsecase) GetPaymentMethod(ctx context.Context, methodName string) (*models.PaymentMethod, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodUsecase.GetPaymentMethod")
	defer sp.Finish()

	tracing.LogRequest(sp, methodName)

	payment, err := u.transactionRepo.FindByMethodName(ctx, methodName)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	return payment, nil
}

func (u *paymentMethodUsecase) GetPaymentMethodById(ctx context.Context, id string) (*models.PaymentMethod, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodUsecase.GetPaymentMethod")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	payment, err := u.transactionRepo.FindByMethodNameById(ctx, id)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	return payment, nil
}
