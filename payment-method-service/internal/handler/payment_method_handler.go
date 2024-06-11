package handler

import (
	"context"

	"payment-method-service/internal/usecase"
	"payment-method-service/lib/tracing"
	"payment-method-service/proto"

	"github.com/opentracing/opentracing-go"
)

type PaymentMethodHandler struct {
	usecase usecase.PaymentMethodUsecase
	proto.UnimplementedPaymentMethodServiceServer
}

func NewPaymentMethodHandler(u usecase.PaymentMethodUsecase) *PaymentMethodHandler {
	return &PaymentMethodHandler{
		usecase: u,
	}
}

func (h *PaymentMethodHandler) CreatePaymentMethod(ctx context.Context, req *proto.CreatePaymentMethodRequest) (*proto.CreatePaymentMethodResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "CreatePaymentMethod")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.CreatePaymentMethod(ctx, req)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &proto.CreatePaymentMethodResponse{
		PaymentMethodId: tx.Id,
		MethodName:      tx.MethodName,
	}, nil
}

func (h *PaymentMethodHandler) GetPaymentMethod(ctx context.Context, req *proto.GetPaymentMethodRequest) (*proto.GetPaymentMethodResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetPaymentMethod")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.GetPaymentMethod(ctx, req.MethodName)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &proto.GetPaymentMethodResponse{
		Id:         tx.Id,
		MethodName: tx.MethodName,
	}, nil
}

func (h *PaymentMethodHandler) GetPaymentMethodById(ctx context.Context, req *proto.GetPaymentMethodByIdRequest) (*proto.GetPaymentMethodResponse, error) {
	sp, ctx := opentracing.StartSpanFromContext(ctx, "GetPaymentMethodById")
	defer sp.Finish()

	tracing.LogRequest(sp, req)

	tx, err := h.usecase.GetPaymentMethodById(ctx, req.Id)
	if err != nil {
		tracing.LogError(sp, err)

		return nil, err
	}

	return &proto.GetPaymentMethodResponse{
		Id:         tx.Id,
		MethodName: tx.MethodName,
	}, nil
}
