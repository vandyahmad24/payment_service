package usecase

import (
	"api-gateway/internal/repository"
	"api-gateway/internal/request"
	"api-gateway/proto"
	"context"
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateTransactionUsecase(t *testing.T) {
	mockRepo := repository.MockTransactionRepository{}
	Convey("Unit Test CreateTransaction", t, func() {

		Convey("Positive Scenario CreateTransaction", func() {

			var req request.CreateTransactionRequest

			var reqRepo proto.CreateTransactionRequest
			var resRepo proto.CreateTransactionResponse
			mockRepo.On("CreateTransaction", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.CreateTransaction(context.Background(), &req)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Negative Scenario CreateTransaction", func() {

			var req request.CreateTransactionRequest

			var reqRepo proto.CreateTransactionRequest
			var resRepo proto.CreateTransactionResponse
			mockRepo.On("CreateTransaction", &reqRepo).Return(&resRepo, errors.New("error")).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.CreateTransaction(context.Background(), &req)
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})

	})
}

func TestGetTransactionUsecase(t *testing.T) {
	mockRepo := repository.MockTransactionRepository{}
	Convey("Unit Test GetTransaction", t, func() {

		Convey("Positive Scenario GetTransaction", func() {

			var reqRepo proto.GetTransactionRequest
			reqRepo.TransactionId = "1"
			var resRepo proto.GetTransactionResponse
			mockRepo.On("GetTransaction", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.GetTransaction(context.Background(), "1")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Negative Scenario GetTransaction", func() {

			var reqRepo proto.GetTransactionRequest
			reqRepo.TransactionId = "1"
			var resRepo proto.GetTransactionResponse
			mockRepo.On("GetTransaction", &reqRepo).Return(&resRepo, errors.New("errors")).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.GetTransaction(context.Background(), "1")
			So(err, ShouldNotBeNil)
			So(resp, ShouldBeNil)
		})
	})
}

func TestGetListTransactionByCustomerIdUsecase(t *testing.T) {
	mockRepo := repository.MockTransactionRepository{}
	Convey("Unit Test GetListTransactionByCustomerId", t, func() {

		Convey("Positive Scenario GetListTransactionByCustomerId", func() {

			var reqRepo proto.ListTransactionsRequest
			reqRepo.CustomerId = "1"
			var resRepo proto.ListTransactionsResponse
			resRepo.Transactions = append(resRepo.Transactions, &proto.GetTransactionResponse{})
			mockRepo.On("ListTransactions", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.GetListTransactionByCustomerId(context.Background(), "1")
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Negative Scenario GetListTransactionByCustomerId", func() {

			var reqRepo proto.ListTransactionsRequest
			reqRepo.CustomerId = "1"
			var resRepo proto.ListTransactionsResponse
			//resRepo.Transactions = append(resRepo.Transactions, &proto.GetTransactionResponse{})
			mockRepo.On("ListTransactions", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.GetListTransactionByCustomerId(context.Background(), "1")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

		Convey("Negative Scenario Failed List", func() {

			var reqRepo proto.ListTransactionsRequest
			reqRepo.CustomerId = "1"
			var resRepo proto.ListTransactionsResponse
			//resRepo.Transactions = append(resRepo.Transactions, &proto.GetTransactionResponse{})
			mockRepo.On("ListTransactions", &reqRepo).Return(&resRepo, errors.New("errors")).Once()
			uc := NewTransactionUsecase(&mockRepo)
			resp, err := uc.GetListTransactionByCustomerId(context.Background(), "1")
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})
	})
}
