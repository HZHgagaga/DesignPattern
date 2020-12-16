package main

import "fmt"

type User struct {
	equip *Equip
	pig   *Pig
	show  *Show
}

//提供了一个统一的接口, 用来访问子系统中的一群接口
func (u *User) Save() {
	u.equip.Save()
	u.pig.Save()
	u.show.Save()
	fmt.Println("User Save...")
}

type Equip struct{}

func (e *Equip) Save() {
	fmt.Println("Equip Save...")
}

type Pig struct{}

func (p *Pig) Save() {
	fmt.Println("Pig Save...")
}

type Show struct{}

func (s *Show) Save() {
	fmt.Println("Show Save...")
}

func main() {
	user := &User{
		equip: &Equip{},
		pig:   &Pig{},
		show:  &Show{},
	}
	user.Save()
}
