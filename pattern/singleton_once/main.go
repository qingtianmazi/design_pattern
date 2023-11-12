package main

import (
	"fmt"
	"sync"
)

// 1.单例类需要是包内私有的，不能被外界访问到，否则就能实例化多个对象
type singletonCar struct {
	name string
}

type SingletonCarInterface interface {
	PrintName()
}

// 2.访问单例对象的指针必须是私有指针，不能被外界访问到，否则外界就能修改这个指针的指向，导致单例对象丢失
var s *singletonCar

func newSingletonCar() *singletonCar {
	return &singletonCar{
		name: "BMW",
	}
}

func (sc *singletonCar) PrintName() {
	fmt.Println(sc.name)
}

var once sync.Once

// 3.使用sync.Once来保证只实例化一次
func GetSingleton() SingletonCarInterface {
	once.Do(func() {
		s = newSingletonCar()
	})
	return s
}

func main() {
	sc := GetSingleton()
	sc.PrintName()
}
