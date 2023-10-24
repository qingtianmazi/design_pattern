package main

import "fmt"

type BMWCar struct {
}

func (c *BMWCar) Run() {
	fmt.Println("BMW car is running")
}

type AudiCar struct {
}

func (c *AudiCar) Run() {
	fmt.Println("Audi car is running")
}

type Zhang3 struct {
}

func (z *Zhang3) DriveBMW(car *BMWCar) {
	fmt.Println("zhang3 is driving car")
	car.Run()
}
func (z *Zhang3) DriveAudo(car *AudiCar) {
	fmt.Println("zhang3 is driving car")
	car.Run()
}

type Li4 struct {
}

func (l *Li4) DriveBMW(car *BMWCar) {
	fmt.Println("li4 is driving car")
	car.Run()
}
func (z *Li4) DriveAudo(car *AudiCar) {
	fmt.Println("Li4 is driving car")
	car.Run()
}

func main() {
	var bmw *BMWCar
	var audi *AudiCar
	var z3 Zhang3
	var l4 Li4
	z3.DriveBMW(bmw)
	z3.DriveAudo(audi)
	l4.DriveBMW(bmw)
	l4.DriveAudo(audi)
}

//// 依赖倒转原则：实现层和业务逻辑层都只依赖于抽象层
//
//// 抽象层
//// 抽象层之间相互依赖
//type Car interface {
//	Run()
//}
//
//type Driver interface {
//	Drive(car Car)
//}
//
//// 实现层
//// 汽车实现层
//type Benz struct {
//}
//
//func (b *Benz) Run() {
//	fmt.Println("benz is running")
//}
//
//type Bmw struct {
//}
//
//func (b *Bmw) Run() {
//	fmt.Println("bmw is running")
//}
//
//// 司机实现层
//type zhang3 struct {
//}
//
//func (z *zhang3) Drive(car Car) {
//	fmt.Println("zhang3 开汽车")
//	car.Run()
//}
//
//type li4 struct {
//}
//
//func (l *li4) Drive(car Car) {
//	fmt.Println("li4 开汽车")
//	car.Run()
//}
//
//// 业务逻辑层
//func main() {
//	// 只依赖于抽象层，针对抽象层编程
//	// 抽象汽车
//	var car Car
//	car = new(Benz)
//	// 抽象司机
//	var driver Driver
//	// 里氏替换原则，用具体实现类替换抽象类
//	driver = new(zhang3)
//	driver.Drive(car) // zhang3 开汽车  benz is running
//	car = new(Bmw)
//	driver = new(li4)
//	driver.Drive(car) // li4 开汽车 bmw is running
//}
