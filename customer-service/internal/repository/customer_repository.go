package repository

import (
	"context"

	"customer-service/internal/models"
	"customer-service/lib/tracing"

	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(ctx context.Context, tx *models.Customer) error
	GetById(ctx context.Context, id string) (*models.Customer, error)
	GetByIds(ctx context.Context, ids []int) ([]*models.Customer, error)
	GetByEmail(ctx context.Context, email string) (*models.Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) GetById(ctx context.Context, id string) (*models.Customer, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerRepository.GetById")
	defer sp.Finish()

	tracing.LogRequest(sp, id)

	var customer models.Customer
	err := r.db.WithContext(ctx).Where("id = ?", id).Find(&customer).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	return &customer, nil

}

func (r *customerRepository) Create(ctx context.Context, tx *models.Customer) error {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerRepository.Create")
	defer sp.Finish()

	tracing.LogRequest(sp, tx)

	return r.db.WithContext(ctx).Create(tx).Error
}

func (r *customerRepository) GetByEmail(ctx context.Context, email string) (*models.Customer, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerRepository.GetByEmail")
	defer sp.Finish()

	tracing.LogRequest(sp, email)

	var customer models.Customer
	err := r.db.WithContext(ctx).Where("email = ?", email).Find(&customer).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	return &customer, nil

}

func (r *customerRepository) GetByIds(ctx context.Context, ids []int) ([]*models.Customer, error) {

	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerRepository.GetByIds")
	defer sp.Finish()

	tracing.LogRequest(sp, ids)

	var customers []*models.Customer
	err := r.db.WithContext(ctx).Where("id in ?", ids).Find(&customers).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	return customers, nil

}
