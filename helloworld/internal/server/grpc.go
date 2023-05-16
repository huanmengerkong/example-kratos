package server

import (
	"context"
	"fmt"
	grpc2 "google.golang.org/grpc"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, greeter *service.GreeterService, user *service.UseSrv, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
		grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc2.UnaryServerInfo, handler grpc2.UnaryHandler) (resp interface{}, err error) {
			fmt.Println(info.FullMethod)
			return handler(ctx, req)
		}),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterGreeterServer(srv, greeter)
	v1.RegisterUserServer(srv, user)
	return srv
}
