package service

import hello "github.com/Cambo9p/sage/services/hello_service"

// a service only needs an entrypoint and a function that will return true when the correct command is called
type Service interface {
	Listen(string)
}

// maybe read from yml file to get the services
func GetAllServices() []Service {
	services := make([]Service, 0)

	h := hello.NewHelloService()
	services = append(services, h)

	return services
}

// // this function will start up all of the services by passing in them the reader
// func startupServices(mqr event.MessageQueueReader) {
// 	services := getAllServices()

// }
