package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"qqbot/api"
	"qqbot/internal/model"
	"time"

	nethttp "net/http"
)

const (
	QQClientAPI = "http://127.0.0.1:3000/"
)

// ProxyService is a proxy of QQ OneBot. It implements all defined api of OneBotV11.
type ProxyService struct {
	api.UnimplementedOneBotServer

	qqClient *http.Client
}

func NewProxyService() *ProxyService {
	qqClient, _ := http.NewClient(context.Background(), http.WithTimeout(1*time.Second))
	return &ProxyService{
		qqClient: qqClient,
	}
}

func (s *ProxyService) SendMsg(ctx context.Context, req *api.SendMsgReq) (resp *api.SendMsgResp, err error) {
	resp = &api.SendMsgResp{}
	defer log.Context(ctx).Infof("[ProxyService] SendMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqCall(ctx, model.EndPointSendMsg, req, resp)
	return
}

func (s *ProxyService) SendPrivateMsg(ctx context.Context, req *api.SendPrivateMsgReq) (resp *api.SendPrivateMsgResp, err error) {
	resp = &api.SendPrivateMsgResp{}
	defer log.Context(ctx).Infof("[ProxyService] SendPrivateMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqCall(ctx, model.EndPointSendPrivateMsg, req, resp)
	return
}

func (s *ProxyService) SendGroupMsg(ctx context.Context, req *api.SendGroupMsgReq) (resp *api.SendGroupMsgResp, err error) {
	resp = &api.SendGroupMsgResp{}
	defer log.Context(ctx).Infof("[ProxyService] SendGroupMsg req: %v, resp: %v, err: %v", req, resp, err)
	err = s.qqCall(ctx, model.EndPointSendGroupMsg, req, resp)
	return
}

func (s *ProxyService) qqCall(ctx context.Context, endpoint string, req interface{}, resp interface{}) (err error) {
	data, err := json.Marshal(req)
	if err != nil {
		return
	}
	hReq, _ := nethttp.NewRequestWithContext(ctx, nethttp.MethodPost, QQClientAPI+endpoint, bytes.NewReader(data))
	hResp, err := s.qqClient.Do(hReq)
	if err != nil {
		return
	}
	defer hResp.Body.Close()
	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(hResp.Body)
	if err != nil {
		return
	}
	qResp := &model.QQCallResp{
		Data: resp,
	}
	err = json.Unmarshal(buffer.Bytes(), qResp)
	if err != nil {
		return
	}
	if qResp.Status != model.QQCallRespStatusOK {
		err = errors.New("qq call resp status not ok")
	}
	return
}
