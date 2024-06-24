package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"qqbot/api"
	"qqbot/internal/model"
	"qqbot/internal/repo"
)

// ProxyService is a proxy of QQ OneBot. It implements all defined api of OneBotV11.
type ProxyService struct {
	api.UnimplementedProxyServer
	qqClient *repo.QQClient
}

func NewProxyService() *ProxyService {
	return &ProxyService{
		qqClient: repo.NewQQClient(),
	}
}

func (s *ProxyService) SendMsg(ctx context.Context, req *api.SendMsgReq) (resp *api.SendMsgResp, err error) {
	resp = &api.SendMsgResp{}
	defer log.Context(ctx).Infof("[proxy] SendMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqClient.Call(ctx, model.EndPointSendMsg, req, resp)
	return
}

func (s *ProxyService) SendPrivateMsg(ctx context.Context, req *api.SendPrivateMsgReq) (resp *api.SendPrivateMsgResp, err error) {
	resp = &api.SendPrivateMsgResp{}
	defer log.Context(ctx).Infof("[proxy] SendPrivateMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqClient.Call(ctx, model.EndPointSendPrivateMsg, req, resp)
	return
}

func (s *ProxyService) SendGroupMsg(ctx context.Context, req *api.SendGroupMsgReq) (resp *api.SendGroupMsgResp, err error) {
	resp = &api.SendGroupMsgResp{}
	defer log.Context(ctx).Infof("[proxy] SendGroupMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqClient.Call(ctx, model.EndPointSendGroupMsg, req, resp)
	return
}

func (s *ProxyService) SendDebugMsg(ctx context.Context, req *api.SendDebugMsgReq) (resp *api.SendDebugMsgResp, err error) {
	resp = &api.SendDebugMsgResp{}
	defer log.Context(ctx).Infof("[proxy] SendDebugMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqClient.Call(ctx, model.EndPointSendGroupMsg, &api.SendGroupMsgReq{
		GroupId:    DebugGroupId,
		Message:    req.GetMessage(),
		AutoEscape: false,
	}, resp)
	return
}
