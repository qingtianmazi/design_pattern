package main

import "fmt"

// 简单工厂模式：对于产品进行抽象，对于工厂没有做抽象

// 抽象层
type Fruit interface {
	Show()
}

// 实现层
type Apple struct {
}

func (a *Apple) Show() {
	fmt.Println("我是apple")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("我是banana")
}

// 此种工厂模式不符合开闭原则
type Factory struct {
}

// 如果新增一个产品对象，则需要修改Factory类的源代码
func (f *Factory) NewFruitFactory(name string) Fruit {
	var fruit Fruit
	if name == "apple" {
		fruit = new(Apple)
	} else if name == "banana" {
		fruit = new(Banana)
	}
	return fruit
}

// 业务逻辑层
func main() {
	var fruit Fruit
	var fac Factory
	fruit = fac.NewFruitFactory("apple")
	fruit.Show()
	fruit = fac.NewFruitFactory("banana")
	fruit.Show()
}
