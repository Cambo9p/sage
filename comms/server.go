package comms

import (
	"bufio"
	"fmt"
	"net"

	"go.uber.org/zap"

	"github.com/Cambo9p/sage/pipeline"
	logging "github.com/Cambo9p/sage/util/logger"
)

var logger *zap.Logger

// gets called when there is a voice connection -- when the server recieves a package from the client
func handleConnection(c net.Conn, pipe pipeline.Pipeline) {
	logger.Info(fmt.Sprintf("Serving %s\n", c.RemoteAddr().String()))

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		// read a line which should be one or more words

		line := scanner.Text()
		fmt.Printf(line)

		pipe.AddMessageToPipeline(line)
	}
}

// TODO refactor to smaller methods?
// TODO remove pipeline dependancy since we will communicate 
// over a channel with the filter
func StartServer(pipe pipeline.Pipeline) {
	logger = logging.NewLogger("comms")
	logger.Info("starting server")

	newConfig, err := getConfig()
	port := fmt.Sprintf(":%d", newConfig.port)
	host := newConfig.host
	logger.Info(fmt.Sprintf("running server on port %s on %s", port, host))

	l, err := net.Listen("tcp4", ":5000")
	if err != nil {
		logger.Error(fmt.Sprintf("failed to listen for connections on %s", port))
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			logger.Error("failed to accept connections on")
		}
		go handleConnection(c, pipe)
	}

}
