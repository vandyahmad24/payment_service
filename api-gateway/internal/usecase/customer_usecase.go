package usecase

import (
	"api-gateway/internal/repository"
	"api-gateway/internal/request"
	"api-gateway/internal/response"
	"api-gateway/lib/tracing"
	"api-gateway/proto"
	"context"

	"github.com/opentracing/opentracing-go"
)

type CustomerUsecase interface {
	LoginCustomer(ctx context.Context, req *request.LoginRequest) (*response.LoginCustomerResponse, error)
	RegisterCustomer(ctx context.Context, req *request.RegisterRequest) (*response.RegisterCustomerResponse, error)
}

type customerUsecase struct {
	customerRepo repository.CustomerRepository
}

func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{
		customerRepo: repo,
	}
}

func (u *customerUsecase) LoginCustomer(ctx context.Context, req *request.LoginRequest) (*response.LoginCustomerResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "LoginCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	resp, err := u.customerRepo.LoginCustomer(ctx, &proto.LoginCustomerRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	response := &response.LoginCustomerResponse{
		Token: resp.Token,
	}

	return response, nil
}

func (u *customerUsecase) RegisterCustomer(ctx context.Context, req *request.RegisterRequest) (*response.RegisterCustomerResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "RegisterCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	resp, err := u.customerRepo.RegisterCustomer(ctx, &proto.CreateCustomerRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	response := &response.RegisterCustomerResponse{
		CustomerId: resp.CustomerId,
	}

	return response, nil
}
