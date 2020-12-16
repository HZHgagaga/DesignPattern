package main

import "fmt"

type Observer interface {
	Update(string)
}

type Manager struct {
	user []Observer
}

func NewManager() *Manager {
	return &Manager{
		user: make([]Observer, 0),
	}
}

//用户注册，观察manager发布的消息
func (m *Manager) Watch(name string, ob Observer) {
	m.user = append(m.user, ob)
	fmt.Println(name, "watch manager")
}

//manager变化通知所有观察者
func (m *Manager) Change(str string) {
	fmt.Println("manager change:", str)
	m.notify(str)
}

func (m *Manager) notify(str string) {
	for _, v := range m.user {
		v.Update(str)
	}
}

type User1 struct {
	name string
}

func (u *User1) Update(str string) {
	fmt.Println("User1 update:", str)
}

type User2 struct {
	name string
}

func (u *User2) Update(str string) {
	fmt.Println("User2 update:", str)
}

type User3 struct {
	name string
}

func (u *User3) Update(str string) {
	fmt.Println("User3 update:", str)
}

func main() {
	manager := NewManager()
	manager.Watch("User1", &User1{"User1"})
	manager.Watch("User2", &User2{"User2"})
	manager.Watch("User3", &User3{"User3"})
	manager.Change("gogogo")
	manager.Change("eateateat")
}
