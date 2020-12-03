package main

import "fmt"

type Server interface {
	Send()
}

type TcpServer struct{}

func (t *TcpServer) Send() {
	fmt.Println("TcpServer send")
}

type HttpServer struct{}

func (h *HttpServer) Send() {
	fmt.Println("HttpServer send")
}

type ServerFactory interface {
	Create() Server
}

type TcpFactory struct{}

func NewTcpFactory() ServerFactory {
	return &TcpFactory{}
}

func (t *TcpFactory) Create() Server {
	return &TcpServer{}
}

type HttpFactory struct{}

func NewHttpFactory() ServerFactory {
	return &HttpFactory{}
}

func (h *HttpFactory) Create() Server {
	return &HttpServer{}
}

func main() {
	factory := NewTcpFactory()
	tcpServer := factory.Create()
	tcpServer.Send()
	factory = NewHttpFactory()
	httpServer := factory.Create()
	httpServer.Send()
}
