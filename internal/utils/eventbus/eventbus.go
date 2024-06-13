package eventbus

import (
	"qqbot/api"
	"qqbot/internal/model"
)

var bus map[api.EventId][]*model.EventChan

func init() {
	bus = make(map[api.EventId][]*model.EventChan)
}

func Sub(ch *model.EventChan, list []api.EventId) <-chan *api.NotifyEvent {
	for _, id := range list {
		switch id {
		case api.EventId_MessageEventAll:
			bus[api.EventId_MessageEventPrivateMessage] = append(bus[api.EventId_MessageEventPrivateMessage], ch)
			bus[api.EventId_MessageEventGroupMessage] = append(bus[api.EventId_MessageEventGroupMessage], ch)
		default:
			bus[id] = append(bus[id], ch)
		}
	}
	return ch.Sub()
}

func Pub(event *api.NotifyEvent) {
	for _, ch := range bus[event.GetEventId()] {
		if !ch.IsClose() {
			ch.Pub(event)
		}
	}
}
