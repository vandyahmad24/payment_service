package handler

import (
	"context"
	"net/http"
	"time"

	"api-gateway/internal/request"
	"api-gateway/internal/usecase"
	"api-gateway/lib"
	"api-gateway/lib/tracing"
	"api-gateway/lib/validate"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

type CustomerHandler struct {
	customerUsecase usecase.CustomerUsecase
}

func NewCustomerHandler(customerUsecase usecase.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customerUsecase,
	}
}

func (h *CustomerHandler) LoginCustomer(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sp, ctx := opentracing.StartSpanFromContext(ctx, "LoginCustomer")
	defer sp.Finish()

	var req request.LoginRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := validate.FormatValidationError(err)
		tracing.LogError(sp, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tracing.LogRequest(sp, req)

	response, err := h.customerUsecase.LoginCustomer(ctx, &req)
	if err != nil {
		tracing.LogError(sp, err)
		respError, code := lib.TranslateErrorToResponse(err)
		c.JSON(code, respError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sp, ctx := opentracing.StartSpanFromContext(ctx, "RegisterCustomer")
	defer sp.Finish()

	var req request.RegisterRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := validate.FormatValidationError(err)
		tracing.LogError(sp, err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	tracing.LogRequest(sp, req)

	response, err := h.customerUsecase.RegisterCustomer(ctx, &req)
	if err != nil {
		tracing.LogError(sp, err)
		respError, code := lib.TranslateErrorToResponse(err)
		c.JSON(code, respError)
		return
	}

	c.JSON(http.StatusOK, response)
}
