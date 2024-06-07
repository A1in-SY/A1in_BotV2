package server

import (
	"qqbot/api"
	"qqbot/internal/conf"
	"qqbot/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewProxyServer(c *conf.Server, proxy *service.ProxyService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Proxy.Network != "" {
		opts = append(opts, grpc.Network(c.Proxy.Network))
	}
	if c.Proxy.Addr != "" {
		opts = append(opts, grpc.Address(c.Proxy.Addr))
	}
	if c.Proxy.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Proxy.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	api.RegisterProxyServer(srv, proxy)
	return srv
}
