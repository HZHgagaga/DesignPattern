package main

import (
	"fmt"
	"sync"
)

type Request interface {
	Process()
}

type User struct {
	name string
}

type RunRequest struct {
	user *User
	data string
}

//封装行为请求
func (a *RunRequest) init(u *User, data string) {
	a.user = u
	a.data = data
	fmt.Println("package RunRequest")
}

func (a *RunRequest) Process() {
	fmt.Println(a.user.name, "run", a.data)
}

type WalkRequest struct {
	user *User
	data string
}

//封装行为请求
func (w *WalkRequest) init(u *User, data string) {
	w.user = u
	w.data = data
	fmt.Println("package walkRequest")
}

func (a *WalkRequest) Process() {
	fmt.Println(a.user.name, "walk", a.data)
}

type Actor struct {
	queue chan Request
}

//请求通过channel传递，与动作的请求者解耦
func NewActor() *Actor {
	return &Actor{
		queue: make(chan Request, 10),
	}
}

//把方法封装成request后打入执行者的channel
func (a *Actor) AddMessage(req Request) {
	a.queue <- req
}

func (a *Actor) Run() {
	go func() {
		for {
			select {
			case req := <-a.queue:
				if req == nil {
					wg.Done()
					return
				}
				//不同类型的请求调用相应的行为
				req.Process()
			}
		}
	}()
}

var wg = sync.WaitGroup{}

func main() {
	wg.Add(1)
	user := &User{"hzh"}
	runRequest := new(RunRequest)
	//将一个请求封装为一个对象
	runRequest.init(user, "test")
	walkRequest := new(WalkRequest)
	//将一个请求封装为一个对象
	walkRequest.init(user, "test")

	actor := NewActor()
	//actor持续处理请求
	actor.Run()
	//将请求封装好作为参数传入actor
	fmt.Println("actor add RunRequest")
	actor.AddMessage(runRequest)
	//将请求封装好作为参数传入actor
	fmt.Println("actor add WalkRequest")
	actor.AddMessage(walkRequest)
	actor.AddMessage(nil)
	wg.Wait()
}
