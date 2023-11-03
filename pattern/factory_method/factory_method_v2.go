package main

import "fmt"

// 工厂方法就是：简单工厂+开闭原则

// 抽象层
type Milk interface {
	Brand()
}

type AbstractMilkFactory interface {
	ProduceMilk() Milk
}

// 实现层
// 牛奶类的具体对象
type MengNiu struct {
	brand string
}

func (m *MengNiu) Brand() {
	fmt.Println("品牌:" + m.brand)
}

type YiLi struct {
	brand string
}

func (y *YiLi) Brand() {
	fmt.Println("品牌:" + y.brand)
}

// 工厂类的具体对象
type MengNiuFactory struct {
}

func (m *MengNiuFactory) ProduceMilk() Milk {
	var mengNiu MengNiu
	mengNiu.brand = "蒙牛"
	return &mengNiu
}

type YiLiFactory struct {
}

func (y *YiLiFactory) ProduceMilk() Milk {
	var yili YiLi
	yili.brand = "伊利"
	return &yili
}

// 业务逻辑层
func main() {
	// 生产蒙牛牛奶
	// 1.创建抽象牛奶工厂
	var mnFacroty AbstractMilkFactory
	// 2.实例化成蒙牛工厂
	mnFacroty = new(MengNiuFactory)
	// 3.生产蒙牛牛奶
	mnMilk := mnFacroty.ProduceMilk()
	// 4.显示牛奶品牌
	mnMilk.Brand()
	var ylFactory AbstractMilkFactory
	ylFactory = new(YiLiFactory)
	ylMilk := ylFactory.ProduceMilk()
	ylMilk.Brand()
}
