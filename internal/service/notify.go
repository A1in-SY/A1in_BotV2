package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/websocket"
	nethttp "net/http"
	"qqbot/api"
	"qqbot/internal/repo"
)

type NotifyService struct {
	api.UnimplementedNotifyServer

	wsUpGrader *websocket.Upgrader
}

func NewNotifyService() *NotifyService {
	return &NotifyService{
		wsUpGrader: &websocket.Upgrader{
			CheckOrigin: func(r *nethttp.Request) bool {
				// 允许CORS请求
				return true
			},
		},
	}
}

func (s *NotifyService) Link(w http.ResponseWriter, r *http.Request) {
	ws, err := s.wsUpGrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("[notify] ws upgrade err: %v", err)
		return
	}
	nClient := repo.NewNotifyClient(ws)
	go nClient.Handle()
	return
}
