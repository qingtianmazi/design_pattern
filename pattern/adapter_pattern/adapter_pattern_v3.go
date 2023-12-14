package main

import (
	"fmt"
)

// iPhone手机需要使用电源线
type Lighting interface {
	UsingLighting()
}

// iPhone原装电源线   可以直接用
type LineLighting struct {
	name string
}

func (l *LineLighting) UsingLighting() {
	fmt.Println("我是" + l.name)
}

// type-c电源线   需要适配器转换转接口
type LineTypeC struct {
	name string
}

func (l *LineTypeC) UsingTypeC() {
	fmt.Println("我是" + l.name)
}

// type-c转lighting的适配器
type AdapterTypeC struct {
	lc LineTypeC
}

func (a *AdapterTypeC) UsingLighting() {
	a.lc.UsingTypeC()
	fmt.Println("使用了适配器的方法，变成了lighting")
}

type IPhone struct {
	lineLighting Lighting
}

func (i *IPhone) charge() {
	i.lineLighting.UsingLighting()
}

func main() {
	// 一根typec的数据线
	var typec LineTypeC
	typec = LineTypeC{"type-c"}
	// typec的适配器，适配成lighting
	var ac AdapterTypeC
	ac = AdapterTypeC{typec}
	// 定义一个iphone，需要使用lighting充电
	var iphone IPhone
	iphone = IPhone{&ac}
	iphone.charge()
}
