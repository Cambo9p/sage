package event

// the main idea of this package is a message queue struct (dependancy injection for later testing) that will hold the current
// command events, for example if a user
import (
	"go.uber.org/zap"
)

var logger *zap.Logger

type MessageQueue interface {
	AddMessagetoQueue(message string)
	RecieveMessage() string
}

func NewMessageQueue() MessageQueue {
	return &messageQueueImpl{
		ch: make(chan string),
	}
}

type messageQueueImpl struct {
	ch chan string
}

func (m *messageQueueImpl) AddMessagetoQueue(message string) {
	m.ch <- message
}

func (m *messageQueueImpl) RecieveMessage() string {
	return <-m.ch
}
