package event

// the main idea of this package is a message queue struct (dependancy injection for later testing) that will hold the current
// command events.
// the message queue impl usues dependancy injection so that we split up the reading and writing methods as theyre never used in the same instance

import (
	"go.uber.org/zap"
)

var logger *zap.Logger

type MessageQueueReader interface {
	RecieveMessage() string
}

type MessageQueueWriter interface {
	AddMessagetoQueue(message string)
}

func NewMessageQueueReader() MessageQueueReader {
	return &messageQueueImpl{
		ch: make(chan string),
	}
}

func NewMessageQueueWriter() MessageQueueReader {
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
