package main

import "fmt"

//抽象状态接口
type State interface {
	Live()
}

//环境类、状态上下文
type Human struct {
	state State //人类当前状态，用组合解耦
	time  int
}

//活time个小时，显示当前状态
func (h *Human) Live(time int) {
	h.time += time
	h.state.Live()
}

func (h *Human) SetState(state State) {
	h.state = state
}

func (h *Human) GetTime() int {
	return h.time
}

func (h *Human) Reset() {
	if h.time > 24 {
		h.time = h.time - 24
	}
}

//具体状态，不同状态不同定义
type Sleep struct {
	human *Human
}

//进行行为，可能改变状态
func (s *Sleep) Live() {
	s.human.Reset()
	if s.human.GetTime() > 7 {
		//当前状态结束，改变状态
		s.human.SetState(&GetUp{s.human})
		s.human.Live(0)
	} else {
		fmt.Println("sleep...")
	}
}

//具体状态，不同状态不同定义
type GetUp struct {
	human *Human
}

//进行行为，可能改变状态
func (g *GetUp) Live() {
	if g.human.GetTime() > 9 {
		//当前状态结束，改变状态
		g.human.SetState(&Work{g.human})
		g.human.Live(0)
	} else {
		fmt.Println("get up...")
	}
}

//具体状态，不同状态不同定义
type Work struct {
	human *Human
}

//进行行为，可能改变状态
func (w *Work) Live() {
	if w.human.GetTime() > 21 {
		//当前状态结束，改变状态
		w.human.SetState(&GetOff{w.human})
		w.human.Live(0)
	} else {
		fmt.Println("work...")
	}
}

//不同状态不同定义
type GetOff struct {
	human *Human
}

//进行行为，可能改变状态
func (g *GetOff) Live() {
	if g.human.GetTime() > 25 {
		//当前状态结束，改变状态
		g.human.SetState(&Sleep{g.human})
		g.human.Live(0)
	} else {
		fmt.Println("get off...")
	}
}

func main() {
	human := &Human{
		time: 3,
	}
	human.SetState(&Sleep{human})
	//生活2个小时且进行行为
	human.Live(2)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
	human.Live(1)
}
