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

func NewServer(typ string) Server {
	switch typ {
	case "tcp":
		return &TcpServer{}
	case "http":
		return &HttpServer{}
	}
	return nil
}

func main() {
	simpleServer := NewServer("tcp")
	simpleServer.Send()
	simpleServer = NewServer("http")
	simpleServer.Send()
}
