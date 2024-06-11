package lib

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func WrapErrorWithCode(err error, code codes.Code) error {
	return status.Errorf(code, err.Error())
}

var (
	ErrInvalidArgument  = status.Error(codes.InvalidArgument, "Invalid argument")
	ErrNotFound         = status.Error(codes.NotFound, "Resource not found")
	ErrInternal         = status.Error(codes.Internal, "Internal server error")
	ErrPermissionDenied = status.Error(codes.PermissionDenied, "Permission denied")
	ErrAlreadyExists    = status.Error(codes.AlreadyExists, "User already exists")
)
