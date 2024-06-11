package handler

import (
	"context"
	"time"

	"customer-service/internal/usecase"
	"customer-service/lib/tracing"
	customer "customer-service/proto"

	"github.com/opentracing/opentracing-go"
)

type CustomerHandler struct {
	usecase usecase.CustomerUsecase
	customer.UnimplementedCustomerServiceServer
}

func NewCustomerHandler(u usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		usecase: u,
	}
}

func (h *CustomerHandler) RegisterCustomer(ctx context.Context, req *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "RegisterCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.RegisterCustomer(ctx, req)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &customer.CreateCustomerResponse{
		CustomerId: tx.Id,
		CreatedAt:  tx.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (h *CustomerHandler) LoginCustomer(ctx context.Context, req *customer.LoginCustomerRequest) (*customer.LoginCustomerResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "LoginCustomer")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	stringToken, err := h.usecase.LoginCustomer(ctx, req)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &customer.LoginCustomerResponse{
		Token: stringToken,
	}, nil

}
