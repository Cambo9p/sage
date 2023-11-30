package service

// a service only needs an entrypoint and a function that will return true when the correct command is called
type Service interface {
	Listen(string)
	execute([]string)
}

// maybe read from yml file to get the services
// func getAllServices() []Service {

// }

// // this function will start up all of the services by passing in them the reader
// func startupServices(mqr event.MessageQueueReader) {
// 	services := getAllServices()

// }
