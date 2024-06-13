package model

import "qqbot/api"

type EventChan struct {
	ch      chan *api.NotifyEvent
	isClose bool
}

func NewEventChan(ch chan *api.NotifyEvent) *EventChan {
	return &EventChan{
		ch:      ch,
		isClose: false,
	}
}

func (ch *EventChan) IsClose() bool {
	return ch.isClose
}

func (ch *EventChan) Close() {
	ch.isClose = true
	close(ch.ch)
}

func (ch *EventChan) Pub(event *api.NotifyEvent) {
	ch.ch <- event
}

func (ch *EventChan) Sub() <-chan *api.NotifyEvent {
	return ch.ch
}
