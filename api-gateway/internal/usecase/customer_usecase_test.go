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

func TestLoginCustomerUsecase(t *testing.T) {
	mockRepo := repository.MockCustomerRepository{}
	Convey("Unit Test LoginCustomer", t, func() {

		Convey("Positive Scenario Login", func() {

			var req request.LoginRequest

			var reqRepo proto.LoginCustomerRequest
			var resRepo proto.LoginCustomerResponse
			mockRepo.On("LoginCustomer", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewCustomerUsecase(&mockRepo)
			resp, err := uc.LoginCustomer(context.Background(), &req)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Negative Scenario Login", func() {

			var req request.LoginRequest

			var reqRepo proto.LoginCustomerRequest
			var resRepo proto.LoginCustomerResponse
			mockRepo.On("LoginCustomer", &reqRepo).Return(&resRepo, errors.New("error")).Once()
			uc := NewCustomerUsecase(&mockRepo)
			resp, err := uc.LoginCustomer(context.Background(), &req)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestRegisterCustomerUsecase(t *testing.T) {
	mockRepo := repository.MockCustomerRepository{}
	Convey("Unit Test RegisterCustomer", t, func() {

		Convey("Positive Scenario RegisterCustomer", func() {

			var req request.RegisterRequest

			var reqRepo proto.CreateCustomerRequest
			var resRepo proto.CreateCustomerResponse
			mockRepo.On("RegisterCustomer", &reqRepo).Return(&resRepo, nil).Once()
			uc := NewCustomerUsecase(&mockRepo)
			resp, err := uc.RegisterCustomer(context.Background(), &req)
			So(err, ShouldBeNil)
			So(resp, ShouldNotBeNil)
		})

		Convey("Negative Scenario RegisterCustomer", func() {

			var req request.RegisterRequest

			var reqRepo proto.CreateCustomerRequest
			var resRepo proto.CreateCustomerResponse
			mockRepo.On("RegisterCustomer", &reqRepo).Return(&resRepo, errors.New("error")).Once()
			uc := NewCustomerUsecase(&mockRepo)
			resp, err := uc.RegisterCustomer(context.Background(), &req)
			So(resp, ShouldBeNil)
			So(err, ShouldNotBeNil)
		})

	})
}
