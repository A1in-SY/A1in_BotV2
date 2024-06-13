package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"qqbot/api"
	"qqbot/internal/model"
	"qqbot/internal/utils/eventbus"
)

// ReportService receive event report from QQ.
type ReportService struct {
	api.UnimplementedProxyServer
}

func NewReportService() *ReportService {
	return &ReportService{}
}

func (s *ReportService) ReportEvent(ctx context.Context, req *api.ReportEventReq) (resp *api.ReportEventResp, err error) {
	switch req.GetPostType() {
	case model.PostTypeMessage:
		return s.handleMessage(ctx, req)
	default:
		return nil, fmt.Errorf("unknown post_type")
	}
}

func (s *ReportService) handleMessage(ctx context.Context, req *api.ReportEventReq) (resp *api.ReportEventResp, err error) {
	switch req.GetMessageType() {
	case model.MessageTypeGroup:
		return s.handleGroupMsg(ctx, req)
	default:
		return nil, fmt.Errorf("unknown message_type")
	}
}

func (s *ReportService) handleGroupMsg(ctx context.Context, req *api.ReportEventReq) (resp *api.ReportEventResp, err error) {
	resp = &api.ReportEventResp{}
	defer log.Context(ctx).Infof("[report] handleGroupMsg req: %v, resp: %v, err: %v", req, resp, err)
	data := &api.NotifyEvent_GroupMsg{
		GroupMsg: &api.GroupMessageEvent{
			MessageType: req.GetMessageType(),
			SubType:     req.GetSubType(),
			MessageId:   req.GetMessageId(),
			GroupId:     req.GetGroupId(),
			UserId:      req.GetUserId(),
			Anonymous:   nil,
			Message:     req.GetMessage(),
			RawMessage:  req.GetRawMessage(),
			Font:        req.GetFont(),
			Sender: &api.GroupMessageEvent_Sender{
				UserId:   req.GetSender().GetUserId(),
				Nickname: req.GetSender().GetNickname(),
				Sex:      req.GetSender().GetSex(),
				Age:      req.GetSender().GetAge(),
				Card:     req.GetSender().GetCard(),
				Area:     req.GetSender().GetArea(),
				Level:    req.GetSender().GetLevel(),
				Role:     req.GetSender().GetRole(),
				Title:    req.GetSender().GetTitle(),
			},
		}}
	if req.GetAnonymous() != nil {
		data.GroupMsg.Anonymous = &api.GroupMessageEvent_Anonymous{
			Id:   req.GetAnonymous().GetId(),
			Name: req.GetAnonymous().GetName(),
			Flag: req.GetAnonymous().GetFlag(),
		}
	}
	event := &api.NotifyEvent{
		EventId:         api.EventId_MessageEventGroupMessage,
		Time:            req.GetTime(),
		SelfId:          req.GetSelfId(),
		PostType:        req.GetPostType(),
		NotifyEventData: data,
	}
	eventbus.Pub(event)
	return
}
