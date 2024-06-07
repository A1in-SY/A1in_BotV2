// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"qqbot/internal/conf"
	"qqbot/internal/server"
	"qqbot/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	proxyService := service.NewProxyService()
	grpcServer := server.NewProxyServer(confServer, proxyService, logger)
	notifyService := service.NewNotifyService()
	httpServer1 := server.NewNotifyServer(confServer, notifyService, logger)
	reportService := service.NewReportService()
	httpServer2 := server.NewReportServer(confServer, reportService, logger)
	app := newApp(logger, grpcServer, httpServer1, httpServer2)
	return app, func() {
	}, nil
}
