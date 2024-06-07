package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	nethttp "net/http"
	"qqbot/api"
)

type NotifyService struct {
	api.UnimplementedNotifyServer

	wsUpgrader *websocket.Upgrader
}

func NewNotifyService() *NotifyService {
	return &NotifyService{
		wsUpgrader: &websocket.Upgrader{
			CheckOrigin: func(r *nethttp.Request) bool {
				// 允许CORS请求
				return true
			},
		},
	}
}

func (s *NotifyService) Link(w http.ResponseWriter, r *http.Request) {
	ws, err := s.wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("[notify] ws upgrade err: %v", err)
		return
	}
	go s.handle(ws)
	return
}

func (s *NotifyService) handle(ws *websocket.Conn) {
	defer ws.Close()
	for {
		messageType, p, err := ws.ReadMessage()
		if err != nil {
			log.Errorf("[notify] ws read message err: %v", err)
			break
		}
		log.Infof("[notify] ws receive message: %v, type: %v", string(p), messageType)
	}
}
