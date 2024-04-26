package filter

import (
	"go.uber.org/zap"
	pipeline "github.com/Cambo9p/sage/pipeline"
    logging "github.com/Cambo9p/sage/util/logger"
)

var logger *zap.Logger

// the Filter will recieve the incoming messages
// and listen for the keyword, then it will wait
// to recieve more context before sending it to
// the pipeline
type Filter interface {
    Start()
}
	
// filterImpl takes a channel to recieve 
// messages from the server
// and a pipeline to send commands to the
// services
type filterImpl struct {
    pl pipeline.Pipeline
    c chan string
}

func NewFilter(c chan string, pl pipeline.Pipeline) Filter {
	return &filterImpl{
        c: c, 
        pl: pl,
	}
}

// the server will send its message to the via a channel
func (m *filterImpl) Start() {
	logger = logging.NewLogger("comms")
	logger.Info("starting filter")
    // 1) wait until the key is Recieved
    // 2) buffer till we have enough context
}
