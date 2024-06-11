package repository

import (
	"api-gateway/lib/tracing"
	proto "api-gateway/proto"
	"context"

	"github.com/opentracing/opentracing-go"
)

type CustomerRepository interface {
	LoginCustomer(ctx context.Context, req *proto.LoginCustomerRequest) (*proto.LoginCustomerResponse, error)
	RegisterCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error)
}

type customerRepository struct {
	client proto.CustomerServiceClient
}

func NewCustomerRepository(client proto.CustomerServiceClient) CustomerRepository {
	return &customerRepository{
		client: client,
	}
}

func (r *customerRepository) LoginCustomer(ctx context.Context, req *proto.LoginCustomerRequest) (*proto.LoginCustomerResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "LoginCustomer")
	defer span.Finish()

	tracing.LogRequest(span, req)

	resp, err := r.client.LoginCustomer(ctx, &proto.LoginCustomerRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return resp, nil
}

func (r *customerRepository) RegisterCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RegisterCustomer")
	defer span.Finish()

	tracing.LogRequest(span, req)

	resp, err := r.client.RegisterCustomer(ctx, &proto.CreateCustomerRequest{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	})
	if err != nil {
		tracing.LogError(span, err)
		return nil, err
	}

	return resp, nil
}
