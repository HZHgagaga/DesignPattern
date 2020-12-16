package main

import "fmt"

const (
	MsgEat   = 1
	MsgRun   = 2
	MsgSleep = 3
)

type Router interface {
	Handle()
}

type MsgHandle struct {
	msgMap map[int]Router
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		msgMap: make(map[int]Router),
	}
}

//注册不同的行为
func (m *MsgHandle) Register(id int, router Router) {
	m.msgMap[id] = router
}

//根据行为id进行不同的处理
func (m *MsgHandle) DoMsgHandle(id int) {
	if router, ok := m.msgMap[id]; ok {
		router.Handle()
	} else {
		fmt.Println("find router err")
	}
}

type Eat struct{}

func (e *Eat) Handle() {
	fmt.Println("Eat Eat Eat")
}

type Run struct{}

func (r *Run) Handle() {
	fmt.Println("Run Run Run")
}

type Sleep struct{}

func (r *Sleep) Handle() {
	fmt.Println("Sleep Sleep Sleep")
}

func main() {
	manager := NewMsgHandle()
	manager.Register(MsgEat, &Eat{})
	manager.Register(MsgRun, &Run{})
	manager.Register(MsgSleep, &Sleep{})
	manager.DoMsgHandle(MsgEat)
	manager.DoMsgHandle(MsgRun)
	manager.DoMsgHandle(MsgSleep)
}
