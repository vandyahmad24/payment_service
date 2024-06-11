package repository

import (
	transaction "api-gateway/proto"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockTransactionRepository struct {
	mock.Mock
}

func (m *MockTransactionRepository) CreateTransaction(ctx context.Context, req *transaction.CreateTransactionRequest) (*transaction.CreateTransactionResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*transaction.CreateTransactionResponse), args.Error(1)
}

func (m *MockTransactionRepository) GetTransaction(ctx context.Context, req *transaction.GetTransactionRequest) (*transaction.GetTransactionResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*transaction.GetTransactionResponse), args.Error(1)
}

func (m *MockTransactionRepository) ListTransactions(ctx context.Context, req *transaction.ListTransactionsRequest) (*transaction.ListTransactionsResponse, error) {
	args := m.Called(req)
	return args.Get(0).(*transaction.ListTransactionsResponse), args.Error(1)
}
