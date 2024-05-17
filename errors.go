package gosdk

import (
	"errors"
	"fmt"
	"runtime"

	common "github.com/begonia-org/go-sdk/common/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
)

type SrvError struct {
	Err      error
	ErrCode  int32
	GrpcCode codes.Code
	Action   string
}

func Is(err error, target error) bool {
	return errors.Is(err, target)
}
func As(err error, target interface{}) bool {
	return errors.As(err, target)

}

type Options func(*common.Errors)

func WithClientMessage(msg string) Options {
	return func(e *common.Errors) {
		e.ToClientMessage = msg
	}

}

func NewError(err error, code int32, grpcCode codes.Code, action string, opts ...Options) error {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}
	fn := runtime.FuncForPC(pc)
	funcName := "unknown"
	if fn != nil {
		funcName = fn.Name()
	}

	srvErr := &common.Errors{
		Code:    code,
		Message: err.Error(),
		Action:  action,
		File:    file,
		Line:    int32(line),
		Fn:      funcName,
	}

	for _, opt := range opts {
		opt(srvErr)
	}
	st := status.New(grpcCode, err.Error())
	detailProto, err := anypb.New(srvErr)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to marshal error details: %v", err)
	}
	st, err = st.WithDetails(detailProto)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to marshal error details: %v", err)

	}
	return st.Err()

}
func (s *SrvError) Error() string {
	return fmt.Sprintf("%s|%d", s.Err.Error(), s.ErrCode)
}
func (s *SrvError) Code() int32 {
	return s.ErrCode
}
