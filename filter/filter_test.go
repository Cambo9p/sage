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


func TestFilter_LoneKeyword(t *testing.T) {
    c := make(chan string)
    defer close(c)

    p := &mockPipeline{}

    f := filter.NewFilter(c, p)
    go f.Start()

    go func() {
        c <- "hello" // TODO: use YAML config
    }()

    time.Sleep(100 * time.Millisecond) // account for channel slow

    f.Stop()

    if !p.messageToPipelineCalled {
        t.Errorf("filter_recieve_test() failed to call AddMessageToPipeline()")
    }
}

func TestFilter_NoKeyword(t *testing.T) {
    c := make(chan string)
    defer close(c)

    p := &mockPipeline{}

    f := filter.NewFilter(c, p)
    go f.Start()

    go func() {
        c <- "test"
        c <- "testing"
    }()

    time.Sleep(100 * time.Millisecond) // account for channel slow

    f.Stop()

    if p.messageToPipelineCalled {
        t.Errorf("filter_recieve_test() called AddMessageToPipeline() incorrectly")
    }
}


func TestFilter_KeywordMupliple(t *testing.T) {
    c := make(chan string)
    defer close(c)

    p := &mockPipeline{}

    f := filter.NewFilter(c, p)
    go f.Start()

    go func() {
        c <- "test"
        c <- "testing"
        c <- "hello" // TODO: use YAML config
        c <- "testing"
    }()

    time.Sleep(100 * time.Millisecond) // account for channel slow

    f.Stop()

    if !p.messageToPipelineCalled {
        t.Errorf("filter_recieve_test() failed to call AddMessageToPipeline()")
    }
}
