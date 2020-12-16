package main

import "fmt"

type People interface {
	Description()
}

type Dio struct{}

func (d *Dio) Description() {
	fmt.Println("Dio")
}

type Mask struct {
	people    People //使用组合装饰
	attribute string
}

func (m *Mask) Description() {
	fmt.Print(m.attribute)
	m.people.Description()
}

type Power struct {
	people    People
	attribute string
}

func (p *Power) Description() {
	fmt.Print(p.attribute)
	p.people.Description()
}

func main() {
	dio := &Dio{}
	maskDio := &Mask{
		people:    dio,
		attribute: "不老不死的",
	}
	powerDio := &Power{
		people:    maskDio,
		attribute: "拥有stand power的",
	}
	powerDio.Description()
}
