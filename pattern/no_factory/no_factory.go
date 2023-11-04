package main

import "fmt"

// 不使用工厂方法模式
// 实现层
type Car struct {
	name string
}

func (c *Car) Show(name string) {
	if name == "CHEVROLET" {
		c.name = "保时捷"
	} else if name == "BMW" {
		c.name = "宝马"
	} else {
		c.name = "奥迪"
	}
	fmt.Println("我是：" + c.name)
}

func NewCar(name string) Car {
	var c Car
	if name == "BMW" {
		c.name = "宝马"
	} else if name == "AUDI" {
		c.name = "奥迪"
	} else {
		c.name = "保时捷"
	}
	return c
}

// 业务逻辑层
func main() {
	// 定义Car
	var c Car
	// 创建一个具体的AUDICar
	c = NewCar("AUDI")
	c.Show("AUDI") // 我是：奥迪
	c = NewCar("BMW")
	c.Show("BMW") // 我是：宝马
	c = NewCar("CHEVROLET")
	c.Show("CHEVROLET") // 我是：保时捷
}
