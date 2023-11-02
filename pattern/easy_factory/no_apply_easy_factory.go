package main

import "fmt"

type Car struct {
	name string
}

func NewCar(name string) *Car {
	car := &Car{}
	if name == "BMW" {
		car.name = "BMW"
	} else if name == "Benz" {
		car.name = "Benz"
	} else if name == "Audi" {
		car.name = "Audi"
	}
	return car
}

func (c *Car) Show() {
	fmt.Println("我是：" + c.name)
}

func main() {
	car := NewCar("BMW")
	car.Show()
	car = NewCar("Benz")
	car.Show()
	car = NewCar("Audi")
	car.Show()
}
