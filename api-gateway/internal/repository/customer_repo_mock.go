package repository

import (
	"api-gateway/proto"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockCustomerRepository struct {
	mock.Mock
}

func (m *MockCustomerRepository) LoginCustomer(ctx context.Context, req *proto.LoginCustomerRequest) (*proto.LoginCustomerResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*proto.LoginCustomerResponse), args.Error(1)
}

func (m *MockCustomerRepository) RegisterCustomer(ctx context.Context, req *proto.CreateCustomerRequest) (*proto.CreateCustomerResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*proto.CreateCustomerResponse), args.Error(1)
}
