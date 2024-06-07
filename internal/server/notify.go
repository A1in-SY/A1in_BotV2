package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"qqbot/internal/conf"
	"qqbot/internal/service"
)

// NewNotifyServer new an HTTP server.
func NewNotifyServer(c *conf.Server, notify *service.NotifyService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Notify.Network != "" {
		opts = append(opts, http.Network(c.Notify.Network))
	}
	if c.Notify.Addr != "" {
		opts = append(opts, http.Address(c.Notify.Addr))
	}
	if c.Notify.Timeout != nil {
		opts = append(opts, http.Timeout(c.Notify.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	srv.HandleFunc("/link", notify.Link)
	return srv
}
