package main

import "fmt"

//抽象
type Connection interface {
	ReadMessage()
}

//实现
type Session interface {
	Process()
}

type TcpConnection struct {
	session Session //使用组合解耦
}

//Connection可以独立变化
func (t *TcpConnection) ReadMessage() {
	fmt.Print("TcpConnection read message:")
	t.session.Process()
}

//Session也可以独立变化
type Player struct{}

func (p *Player) Process() {
	fmt.Println("Player process message")
}

//Connection可以独立变化
type UdpConnection struct {
	session Session //使用组合解耦
}

func (u *UdpConnection) ReadMessage() {
	fmt.Print("UdpConnection read message:")
	u.session.Process()
}

//Session也可以独立变化
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
