package helloservice

type HelloService struct {
	message string
}

func newHelloService() *HelloService {
	return &HelloService{
		message: "Hello",
	}
}

func (h *HelloService) listen(message string) bool {
	return message == h.message
}

func (h *HelloService) execute(message []string) bool {
	return message[0] == h.message
}
