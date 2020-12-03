package main

import "fmt"

type Server interface {
	Read()
}

type Client interface {
	Write()
}

type ServerFactory interface {
	CreateServer() Server
	CreateClient() Client
}

type TcpServer struct{}

func (t *TcpServer) Read() {
	fmt.Println("TcpServer read")
}

type TcpClient struct{}

func (t *TcpClient) Write() {
	fmt.Println("TcpClient write")
}

type HttpServer struct{}

func (h *HttpServer) Read() {
	fmt.Println("HttpServer read")
}

type HttpClient struct{}

func (h *HttpClient) Write() {
	fmt.Println("HttpClient write")
}

type TcpFactory struct{}

func (t *TcpFactory) CreateServer() Server {
	return &TcpServer{}
}

func (t *TcpFactory) CreateClient() Client {
	return &TcpClient{}
}

type HttpFactory struct{}

func (t *HttpFactory) CreateServer() Server {
	return &HttpServer{}
}

func (t *HttpFactory) CreateClient() Client {
	return &HttpClient{}
}

func main() {
	tcpFactory := &TcpFactory{}
	tcpClient := tcpFactory.CreateClient()
	fmt.Println("Create TcpClient")
	tcpClient.Write()
	tcpServer := tcpFactory.CreateServer()
	fmt.Println("Create TcpServer")
	tcpServer.Read()

	httpFactory := &HttpFactory{}
	httpClient := httpFactory.CreateClient()
	fmt.Println("Create HttpClient")
	httpClient.Write()
	httpServer := httpFactory.CreateServer()
	fmt.Println("Create HttpServer")
	httpServer.Read()
}
