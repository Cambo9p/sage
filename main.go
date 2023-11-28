package main

import (
	"github.com/Cambo9p/sage/comms"
	"github.com/Cambo9p/sage/event"
)

func main() {
	mq := event.NewMessageQueue()
	comms.StartServer(mq)
}
