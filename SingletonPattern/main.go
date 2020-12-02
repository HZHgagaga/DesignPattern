package main

import (
	"fmt"
	"sync"
)

//懒汉
type Singleton1 struct{}

var instance1 *Singleton1 = &Singleton1{}

func GetSingleton1() *Singleton1 {
	return instance1
}

//饿汉
type Singleton2 struct{}

var instance2 *Singleton2
var mutex sync.Mutex

func GetSingleton2() *Singleton2 {
	if instance2 == nil {
		mutex.Lock()
		defer mutex.Unlock()
		if instance2 == nil {
			fmt.Println("new ins2")
			instance2 = &Singleton2{}
		}
	}
	return instance2
}

//sync.once
type Singleton3 struct{}

var instance3 *Singleton3
var once sync.Once

func GetSingleton3() *Singleton3 {
	once.Do(func() {
		fmt.Println("new ins3")
		instance3 = &Singleton3{}
	})
	return instance3
}

func main() {
	ins1 := GetSingleton1()
	fmt.Printf("%p\n", ins1)
	ins1 = GetSingleton1()
	fmt.Printf("%p\n", ins1)

	ins2 := GetSingleton2()
	fmt.Printf("%p\n", ins2)
	ins2 = GetSingleton2()
	fmt.Printf("%p\n", ins2)

	ins3 := GetSingleton3()
	fmt.Printf("%p\n", ins3)
	ins3 = GetSingleton3()
	fmt.Printf("%p\n", ins3)
}
