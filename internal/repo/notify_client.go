package repo

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"qqbot/api"
	"qqbot/internal/model"
	"qqbot/internal/utils/eventbus"
)

type NotifyClient struct {
	conn      *websocket.Conn
	isAuth    bool
	name      string
	eventChan *model.EventChan
	qqClient  *QQClient
}

func NewNotifyClient(conn *websocket.Conn) *NotifyClient {
	ch := make(chan *api.NotifyEvent)
	eventChan := model.NewEventChan(ch)
	return &NotifyClient{
		conn:      conn,
		isAuth:    false,
		name:      "",
		eventChan: eventChan,
		qqClient:  NewQQClient(),
	}
}

func (c *NotifyClient) Handle() {
	var err error
	defer func() {
		c.conn.Close()
		c.eventChan.Close()
		c.qqClient.Call(context.Background(), model.EndPointSendGroupMsg, &api.SendGroupMsgReq{
			GroupId:    253016320,
			Message:    []*api.Segment{api.BuildTextSegment(fmt.Sprintf("%v 断开链接", c.name))},
			AutoEscape: false,
		}, &api.SendGroupMsgResp{})
	}()
	_, data, err := c.conn.ReadMessage()
	if err != nil {
		log.Errorf("[notify] ws read message err: %v", err)
		return
	}
	reg := &api.RegisterReq{}
	err = proto.Unmarshal(data, reg)
	if err != nil {
		log.Errorf("[notify] unmarshal register message err: %v", err)
		return
	}
	log.Infof("[notify] new register: %v", reg)
	// TODO
	c.isAuth = true
	c.name = reg.GetName()
	respData, _ := proto.Marshal(&api.RegisterResp{Success: true})
	err = c.conn.WriteMessage(websocket.BinaryMessage, respData)
	if err != nil {
		log.Errorf("[notify] write register resp message err: %v", err)
		return
	}
	ch := eventbus.Sub(c.eventChan, reg.GetNeedEventList())
	for {
		event := <-ch
		eventData, err := proto.Marshal(event)
		if err != nil {
			log.Errorf("[notify] marshal event message err: %v", err)
			continue
		}
		err = c.conn.WriteMessage(websocket.BinaryMessage, eventData)
		if err != nil {
			log.Errorf("[notify] write event message err: %v", err)
			break
		}
	}
}
