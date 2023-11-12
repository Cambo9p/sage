package comms

import (
	"bufio"
	"fmt"
	"net"

	"go.uber.org/zap"

	logging "github.com/Cambo9p/sage/util/logger"
)

var logger *zap.Logger

// handles the stream of words
func handleConnection(c net.Conn) {
	logger.Info(fmt.Sprintf("Serving %s\n", c.RemoteAddr().String()))

	scanner := bufio.NewScanner(c)

	for scanner.Scan() {
		// read a line which should be one or more words

		line := scanner.Text()
		fmt.Printf(line)

		// words := strings.Fields(line)

		// check if the key words were said and then create an event

	}

}

func StartServer() {
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
		go handleConnection(c)
	}

}
