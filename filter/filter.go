package filter

import (
	pipeline "github.com/Cambo9p/sage/pipeline"
)

// the Filter will recieve the incoming messages
// and listen for the keyword, then it will wait
// to recieve more context before sending it to
// the pipeline
type Filter interface {
    Recieve(message string)
}
	
type filterImpl struct {
    pl pipeline.Pipeline
}

func NewFilter(p pipeline.Pipeline) Filter {
	return &filterImpl{
        pl: p,
	}
}

// the server will send its message to the filter
func (m *filterImpl) Recieve(message string) {
    // 1) wait until the key is Recieved
    // 2) buffer till we have enough context
}
