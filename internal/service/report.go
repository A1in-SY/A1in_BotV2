package service

import (
	"qqbot/api"
)

// ReportService receive event report from QQ.
type ReportService struct {
	api.UnimplementedProxyServer
}

func NewReportService() *ReportService {
	return &ReportService{}
}
