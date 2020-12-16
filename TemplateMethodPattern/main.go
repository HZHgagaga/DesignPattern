package main

import "fmt"

type Commander interface {
	OnBefore()
	Excute()
	OnAfter()
}

func DoCommand(cmd Commander) {
	cmd.OnBefore()
	cmd.Excute()
	cmd.OnAfter()
}

type Command struct{}

func (u *Command) OnBefore() {
	fmt.Println("Command OnBefore")
}

func (u *Command) OnAfter() {
	fmt.Println("Command OnAfter")
}

type EatCommand struct {
	Command
}

//复用Command类，并且改造Excute方法
func (e *EatCommand) Excute() {
	fmt.Println("Eat Eat Eat")
}

type SleepCommand struct {
	Command
}

func (e *SleepCommand) Excute() {
	fmt.Println("Sleep Sleep Sleep")
}

func main() {
	eatCmd := &EatCommand{}
	sleepCmd := &SleepCommand{}
	DoCommand(eatCmd)
	DoCommand(sleepCmd)
}
