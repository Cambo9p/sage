package filter_test

import (
	"testing"
	"time"

	"github.com/Cambo9p/sage/filter"
	"github.com/Cambo9p/sage/service"
)

type mockPipeline struct {
    messageToPipelineCalled bool
}

func (p *mockPipeline) AddMessageToPipeline(message string) { 
    p.messageToPipelineCalled = true
}


func (p *mockPipeline) AttachServiceReader(ser service.Service) { }


func TestFilter_recieve(t *testing.T) {
    c := make(chan string)
    defer close(c)

    p := &mockPipeline{}

    f := filter.NewFilter(c, p)

    go f.Start()

    go func() {
        c <- "hello"
    }()

    time.Sleep(1000 * time.Millisecond)

    f.Stop()

    if !p.messageToPipelineCalled {
        t.Errorf("filter_recieve_test() failed to call AddMessageToPipeline()")
    }
}
