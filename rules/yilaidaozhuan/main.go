package main

import "fmt"

// 依赖倒转原则：实现层和业务逻辑层都只依赖于抽象层

// 抽象层
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// 实现层
// 汽车实现层
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("benz is running")
}

type Bmw struct {
}

func (b *Bmw) Run() {
	fmt.Println("bmw is running")
}

// 司机实现层
type zhang3 struct {
}

func (z *zhang3) Drive(car Car) {
	fmt.Println("zhang3 开汽车")
	car.Run()
}

type li4 struct {
}

func (l *li4) Drive(car Car) {
	fmt.Println("li4 开汽车")
	car.Run()
}

// 业务逻辑层
func main() {
	var car Car
	car = new(Benz)
	var driver Driver
	driver = new(zhang3)
	driver.Drive(car)
	car = new(Bmw)
	driver = new(li4)
	driver.Drive(car)
}
