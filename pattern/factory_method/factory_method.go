package main

import "fmt"

// 工厂方法模式：简单工厂+开闭原则

// 抽象层
type Phone interface {
	Show()
}
type AbstractFactory interface {
	Create() Phone
}

// 实现层
type IPhone struct {
}

func (i *IPhone) Show() {
	fmt.Println("我是 iphone")
}

type Huawei struct {
}

func (h *Huawei) Show() {
	fmt.Println("我是 华为")
}

type IPhoneFactory struct {
}

func (i *IPhoneFactory) Create() Phone {
	var fac Phone
	fac = new(IPhone)
	return fac
}

type HuaweiFactory struct {
}

func (h *HuaweiFactory) Create() Phone {
	var fac Phone
	fac = new(Huawei)
	return fac
}

// 业务逻辑层
func main() {
	// 生产iphone
	// 生产iphone的工厂
	var iphoneFac AbstractFactory // 面向接口编程：符合依赖倒转原则
	iphoneFac = new(IPhoneFactory)
	iphone := iphoneFac.Create()
	iphone.Show()

	var huaweiFac AbstractFactory
	huaweiFac = new(HuaweiFactory)
	huawei := huaweiFac.Create()
	huawei.Show()
}
