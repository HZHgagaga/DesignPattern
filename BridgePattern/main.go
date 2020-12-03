package main

import "fmt"

type Connection interface {
	ReadMessage()
}

type Session interface {
	Process()
}

type TcpConnection struct {
	session Session
}

func (t *TcpConnection) ReadMessage() {
	fmt.Print("TcpConnection read message:")
	t.session.Process()
}

type Player struct{}

func (p *Player) Process() {
	fmt.Println("Player process message")
}

type UdpConnection struct {
	session Session
}

func (u *UdpConnection) ReadMessage() {
	fmt.Print("UdpConnection read message:")
	u.session.Process()
}

type Master struct{}

func (m *Master) Process() {
	fmt.Println("Master process message")
}

func main() {
	user := &Player{}
	conn := &TcpConnection{user}
	conn.ReadMessage()

	userM := &Master{}
	conn = &TcpConnection{userM}
	conn.ReadMessage()

	user = &Player{}
	connU := &UdpConnection{user}
	connU.ReadMessage()

	userM = &Master{}
	connU = &UdpConnection{userM}
	connU.ReadMessage()
}
