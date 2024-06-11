package repository

import (
	"context"

	"payment-method-service/internal/models"
	"payment-method-service/lib/tracing"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	Create(ctx context.Context, tx *models.PaymentMethod) error
	FindByMethodName(ctx context.Context, methodName string) (*models.PaymentMethod, error)
	FindByMethodNameById(ctx context.Context, id string) (*models.PaymentMethod, error)
}

type paymentMethodRepository struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(db *gorm.DB) PaymentMethodRepository {
	return &paymentMethodRepository{
		db: db,
	}
}

func (r *paymentMethodRepository) Create(ctx context.Context, tx *models.PaymentMethod) error {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodRepository.Create")
	defer sp.Finish()

	tracing.LogRequest(sp, tx)

	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *paymentMethodRepository) FindByMethodName(ctx context.Context, methodName string) (*models.PaymentMethod, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodRepository.FindByMethodName")
	defer sp.Finish()

	tracing.LogRequest(sp, methodName)

	var payment models.PaymentMethod
	err := r.db.WithContext(ctx).Where("method_name = ?", methodName).First(&payment).Error
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *paymentMethodRepository) FindByMethodNameById(ctx context.Context, id string) (*models.PaymentMethod, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "PaymentMethodRepository.FindByMethodNameById")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	var payment models.PaymentMethod
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&payment).Error
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
