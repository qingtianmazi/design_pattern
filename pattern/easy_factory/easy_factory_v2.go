package main

import "fmt"

// 练习题：
// 实现简单工厂模式
// 产品类是电脑

// 抽象层
// 抽象电脑，有一个Show方法
type Computer interface {
	Show()
}

// 实现层
// 具体产品类
type Macbook struct {
	Name string
}

func (m *Macbook) Show() {
	fmt.Println(m.Name)
}

type Lenovo struct {
	Name string
}

func (l *Lenovo) Show() {
	fmt.Println(l.Name)
}

// 工厂类
type EasyFactory struct {
}

// 根据入参，生产具体的产品
func (e *EasyFactory) CreateComputer(name string) Computer {
	var c Computer
	if name == "Macbook" {
		c = &Macbook{Name: name}
	} else if name == "Lenovo" {
		c = &Lenovo{Name: name}
	}
	return c
}

// 对产品类符合开闭原则
// 对于工厂类来说，不符合开闭原则，每新增一个产品种类，都需要修改工厂类的方法
// 业务逻辑
func main() {
	// 1.创建一个抽象产品
	var c Computer
	// 2.创建一个具体工厂对象
	var ef = EasyFactory{}
	// 3.调用工厂方法，生产具体产品
	c = ef.CreateComputer("Macbook")
	// 4.调用具体产品的方法
	c.Show()
	// 5.再次调用工厂方法，生产具体产品 (实现多态)
	c = ef.CreateComputer("Lenovo")
	c.Show()
}
