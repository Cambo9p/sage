package helloservice

import "fmt"

type HelloService struct {
	message string
}

func NewHelloService() *HelloService {
	return &HelloService{
		message: "Hello",
	}
}

func (h *HelloService) Listen(message string) {
	fmt.Printf("you said %s ")
}
