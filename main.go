package main

import (
	"github.com/Cambo9p/sage/comms"
	"github.com/Cambo9p/sage/pipeline"
	"github.com/Cambo9p/sage/service"
)

func startupServices(pipe pipeline.Pipeline) {
	services := service.GetAllServices()
	for _, s := range services {
		pipe.AttachServiceReader(s)
	}
}

func main() {
	pipe := pipeline.NewPipeline()
	comms.StartServer(pipe)

	startupServices(pipe)
}
