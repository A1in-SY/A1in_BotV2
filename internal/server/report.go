package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"qqbot/internal/conf"
	"qqbot/internal/service"
)

func NewReportServer(c *conf.Server, notify *service.ReportService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Report.Network != "" {
		opts = append(opts, http.Network(c.Report.Network))
	}
	if c.Report.Addr != "" {
		opts = append(opts, http.Address(c.Report.Addr))
	}
	if c.Report.Timeout != nil {
		opts = append(opts, http.Timeout(c.Report.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	return srv
}
