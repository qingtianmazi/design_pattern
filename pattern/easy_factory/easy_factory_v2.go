package main

import "fmt"

// 练习题：
// 实现简单工厂模式
// 产品类是电脑

// 抽象层
type Computer interface {
	Show()
}

// 实现层
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

func (e *EasyFactory) CreateComputer(name string) Computer {
	var c Computer
	if name == "Macbook" {
		c = &Macbook{Name: name}
	} else if name == "Lenovo" {
		c = &Lenovo{Name: name}
	}
	return c
}

// 对于工厂类来说，不符合开闭原则，每新增一个产品种类，都需要修改工厂类的方法
// 业务逻辑
func main() {
	var c Computer
	var ef = EasyFactory{}
	c = ef.CreateComputer("Macbook")
	c.Show()
	c = ef.CreateComputer("Lenovo")
	c.Show()
}
