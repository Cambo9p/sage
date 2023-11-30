// pipeline object handles transfer of events
// TODO -- split up reader and writer?
package pipeline

import (
	"sync"

	"github.com/Cambo9p/sage/service"
)

type Pipeline interface {
	AddMessageToPipeline(string)
	AttachServiceReader(service.Service)
}

type pipelineImpl struct {
	channel chan string
	wg      sync.WaitGroup
}

func NewPipeline() Pipeline {
	return &pipelineImpl{
		channel: make(chan string),
	}
}

func (m *pipelineImpl) AddMessageToPipeline(message string) {
	m.channel <- message
}

func (p *pipelineImpl) AttachServiceReader(ser service.Service) {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for message := range p.channel {
			ser.Listen(message)
		}
	}()
}

func (p *pipelineImpl) Close() {
	close(p.channel)
	p.wg.Wait()
}
