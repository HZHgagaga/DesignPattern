package main

import "fmt"

type Connection interface {
	Connect()
	Close()
}

type Session struct {
	addr string
}

func (s *Session) Conn(addr string) {
	fmt.Println("Connect to", addr)
	s.addr = addr
}

func (s *Session) Close() {
	fmt.Println("Close to", s.addr)
	s.addr = ""
}

func (s *Session) Send(str string) {
	fmt.Println(s.addr, "Send:", str)
}

type ProxySession struct {
	session *Session
}

func (p *ProxySession) Wrap(s *Session) {
	p.session = s
}

func (p *ProxySession) Conn(addr string) {
	fmt.Print("经过proxy ")
	p.session.Conn(addr)
}

func (p *ProxySession) Close() {
	fmt.Print("经过proxy ")
	p.session.Close()
}

func (p *ProxySession) Send(str string) {
	fmt.Print("经过proxy ")
	p.session.Send(str)
}

func main() {
	session := &Session{}
	proxySession := &ProxySession{}
	proxySession.Wrap(session)
	proxySession.Conn("127.0.0.1")
	proxySession.Send("hzhgagaga")
	proxySession.Close()
}
