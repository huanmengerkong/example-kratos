package server

import (
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"strconv"
	"user/internal/conf"
	"user/internal/service"
	v1 "user/protogo/adminuser/v1"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
// https://juejin.cn/post/7202409558592782373
func NewGRPCServer(c *conf.Server, user *service.UserService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			selector.Server(user.Server()).Match(user.NewWhiteListMatcher()).Build(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr+":"+strconv.Itoa(int(c.Grpc.Port))))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAdminUserServer(srv, user)

	return srv
}
