package listener

import (
	"io"
)

type listener struct {
	chError chan error

	textMessageListener *WebhookListener
}

func NewMessageListener() MessageListener {
	return &listener{
		chError: make(chan error),
	}
}

func (s *listener) HandleErrors(f func(error)) (closer func()) {
	done := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-s.chError:
				f(err)
			case <-done:
				return
			}

		}
	}()

	return func() {
		done <- struct{}{}
	}
}

type MessageListener interface {
	HandleErrors(f func(error)) (closer func())
	OnMessage(WebhookListener)
	ReadBodyAsync(rawBody io.ReadCloser) error
}

func (s *listener) OnMessage(message WebhookListener) {
	s.textMessageListener = &message
}
