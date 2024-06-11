package lib

import (
	"log"
	"net/http"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func TranslateErrorToResponse(err error) (response ErrorResponse, statusCode int) {
	st, ok := status.FromError(err)
	if !ok {
		return ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Message:    "Internal server error",
		}, 500
	}

	code := st.Code()
	log.Println(code)
	message := st.Message()
	log.Println(message)

	if idx := strings.Index(message, " desc = "); idx != -1 {
		message = message[idx+8:]
	}

	statusCode = mapGRPCCodeToHTTP(code)

	return ErrorResponse{
		StatusCode: statusCode,
		Message:    message,
	}, statusCode
}

func mapGRPCCodeToHTTP(code codes.Code) int {
	switch code {
	case codes.OK:
		return http.StatusOK
	case codes.Canceled:
		return http.StatusRequestTimeout
	case codes.Unknown:
		return http.StatusInternalServerError
	case codes.InvalidArgument:
		return http.StatusBadRequest
	case codes.DeadlineExceeded:
		return http.StatusGatewayTimeout
	case codes.NotFound:
		return http.StatusNotFound
	case codes.AlreadyExists:
		return http.StatusConflict
	case codes.PermissionDenied:
		return http.StatusForbidden
	case codes.ResourceExhausted:
		return http.StatusTooManyRequests
	case codes.FailedPrecondition:
		return http.StatusPreconditionFailed
	case codes.Aborted:
		return http.StatusConflict
	case codes.OutOfRange:
		return http.StatusRequestedRangeNotSatisfiable
	case codes.Unimplemented:
		return http.StatusNotImplemented
	case codes.Internal:
		return http.StatusInternalServerError
	case codes.Unavailable:
		return http.StatusServiceUnavailable
	case codes.DataLoss:
		return http.StatusInternalServerError
	case codes.Unauthenticated:
		return http.StatusUnauthorized
	default:
		return http.StatusInternalServerError
	}
}

func GenerateErrorToResponse(err error, code int) (response ErrorResponse, statusCode int) {

	return ErrorResponse{
		StatusCode: code,
		Message:    err.Error(),
	}, statusCode
}
