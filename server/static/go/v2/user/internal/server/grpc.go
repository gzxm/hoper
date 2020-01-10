package server

import (
	"context"
	"runtime/debug"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	model "github.com/liov/hoper/go/v2/protobuf/user"
	"github.com/liov/hoper/go/v2/user/internal/service"
	"github.com/liov/hoper/go/v2/utils/errorcode"
	"github.com/liov/hoper/go/v2/utils/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func filter(
	ctx context.Context, req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.CallTwo.Errorf("%v panic: %v", info, r)
			debug.PrintStack()
			err = errorcode.SysError.Err()
		}
		//不能添加错误处理，除非所有返回的结构相同
		if err != nil {
			if errcode, ok := err.(errorcode.ErrCode); ok {
				err = errcode.Err()
			}
		}
	}()

	return handler(ctx, req)
}

func Grpc() *grpc.Server {
	s := grpc.NewServer(
		//filter应该在最前
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(filter, grpc_validator.UnaryServerInterceptor())),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(grpc_validator.StreamServerInterceptor())),
	)
	model.RegisterUserServiceServer(s, service.UserSvc)
	// Register reflection service on gRPC server.
	reflection.Register(s)
	return s
}
