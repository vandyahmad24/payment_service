package usecase

import (
	"context"
	"errors"
	"log"
	"time"

	"customer-service/internal/models"
	"customer-service/internal/repository"
	"customer-service/lib"
	"customer-service/lib/jwt"
	"customer-service/lib/tracing"
	customer "customer-service/proto"

	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc/codes"
)

type CustomerUsecase interface {
	RegisterCustomer(ctx context.Context, req *customer.CreateCustomerRequest) (*models.Customer, error)
	LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (string, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(tr repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: tr,
	}
}

func (u *customerUsecase) LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (string, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUsecase.LoginCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	customer, err := u.customerRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		tracing.LogError(sp, err)
		return "", lib.WrapErrorWithCode(lib.ErrPermissionDenied, codes.PermissionDenied)
	}

	isValidPassword := lib.ComparePassword(customer.Password, []byte(req.Password))
	if !isValidPassword {
		tracing.LogError(sp, errors.New("invalid password"))
		return "", lib.WrapErrorWithCode(errors.New("invalid password"), codes.PermissionDenied)
	}

	//generate jwt
	token, err := jwt.GenerateToken(customer.Id)
	if err != nil {
		tracing.LogError(sp, err)
		return "", lib.WrapErrorWithCode(err, codes.PermissionDenied)
	}

	log.Println("token", token)

	//parse
	claimbs, err := jwt.ParseToken(token)
	if err != nil {
		tracing.LogError(sp, err)
		return "", lib.WrapErrorWithCode(err, codes.PermissionDenied)
	}

	tracing.LogObject(sp, "claimbs", claimbs)

	return token, nil
}

func (u *customerUsecase) RegisterCustomer(ctx context.Context, req *customer.CreateCustomerRequest) (*models.Customer, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "CustomerUsecase.RegisterCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	//find email first
	existingCustomer, err := u.customerRepo.GetByEmail(ctx, req.Email)
	if err != nil && err != lib.ErrNotFound {
		tracing.LogError(sp, err)
		return nil, lib.WrapErrorWithCode(lib.ErrInternal, codes.Internal)
	}
	log.Println("existingCustomer", existingCustomer)

	if existingCustomer.Id != "" {
		return nil, lib.WrapErrorWithCode(lib.ErrAlreadyExists, codes.AlreadyExists)
	}

	encryptedPassword, err := lib.GenerateHashFromString(req.Password)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, lib.WrapErrorWithCode(lib.ErrInternal, codes.Internal)
	}

	tx := &models.Customer{
		Id:        uuid.New().String(),
		Name:      req.Name,
		Email:     req.Email,
		Password:  encryptedPassword,
		CreatedAt: time.Now(),
	}

	if err := u.customerRepo.Create(ctx, tx); err != nil {
		tracing.LogError(sp, err)
		return nil, lib.WrapErrorWithCode(lib.ErrInternal, codes.Internal)
	}

	return tx, nil
}
