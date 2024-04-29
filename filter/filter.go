package filter

import (
	pipeline "github.com/Cambo9p/sage/pipeline"
)

// the Filter will recieve the incoming messages
// and listen for the keyword, then it will wait
// to recieve more context before sending it to
// the pipeline
type Filter interface {
    Start()
    Stop()
}
	
// filterImpl takes a channel to recieve 
// messages from the server
// and a pipeline to send commands to the
// services
type filterImpl struct {
    pl pipeline.Pipeline
    c chan string
    stop chan struct{} // TODO: is this the best way to do this?
}

func NewFilter(c chan string, pl pipeline.Pipeline) Filter {
	return &filterImpl{
        c: c, 
        pl: pl,
        stop: make(chan struct{}),
	}
}

func (f *filterImpl) parseConnection() {
    message := ""
    // TODO: buffer so that we can get the next words said
    f.pl.AddMessageToPipeline(message)
}

// the server will send its message to the via a channel
// if the stop channel recieves a struct then we exit
func (f *filterImpl) Start() {
    // 1) wait until the key is Recieved
    // 2) buffer till we have enough context

    for {
        select {
        case msg := <- f.c:
            if msg == "hello" { // TODO: use yaml config
                 go f.parseConnection()
            }
        case <- f.stop:
            return
        }
    }
}

func (f *filterImpl) Stop() {
    close(f.stop)
}
