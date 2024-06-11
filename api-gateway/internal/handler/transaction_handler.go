package handler

import (
	"context"
	"log"
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

type TransactionHandler struct {
	transactionUsecase usecase.TransactionUsecase
}

func NewTransactionHandler(transactionUsecase usecase.TransactionUsecase) *TransactionHandler {
	return &TransactionHandler{
		transactionUsecase: transactionUsecase,
	}
}

func (h *TransactionHandler) CreateTransaction(c *gin.Context) {

	var req request.CreateTransactionRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response := validate.FormatValidationError(err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userId, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found"})
		return
	}

	log.Println("userId ", userId.(string))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sp, ctx := opentracing.StartSpanFromContext(ctx, "CreateTransaction")
	defer sp.Finish()

	req.UserId = userId.(string)

	tracing.LogRequest(sp, req)

	response, err := h.transactionUsecase.CreateTransaction(ctx, &req)
	if err != nil {
		tracing.LogError(sp, err)
		respError, code := lib.TranslateErrorToResponse(err)
		c.JSON(code, respError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetDetailTransaction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetDetailTransaction")
	defer sp.Finish()

	id := c.Param("id")

	tracing.LogRequest(sp, id)

	response, err := h.transactionUsecase.GetTransaction(ctx, id)
	if err != nil {
		tracing.LogError(sp, err)
		respError, code := lib.TranslateErrorToResponse(err)
		c.JSON(code, respError)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (h *TransactionHandler) GetListTransactionByCustomerId(c *gin.Context) {
	// customerId
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetDetailTransaction")
	defer sp.Finish()

	id := c.Param("customerId")

	tracing.LogRequest(sp, id)

	response, err := h.transactionUsecase.GetListTransactionByCustomerId(ctx, id)
	if err != nil {
		tracing.LogError(sp, err)
		respError, code := lib.TranslateErrorToResponse(err)
		c.JSON(code, respError)
		return
	}

	c.JSON(http.StatusOK, response)

}
