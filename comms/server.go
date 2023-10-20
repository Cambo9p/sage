package comms

import "github.com/Cambo9p/sage/util/logger"

func StartServer() error {
  logger, err := logger.NewLogger("comms") 
  if err != nil {
    return err
  }

  logger.Info("starting server")

  return nil
}
