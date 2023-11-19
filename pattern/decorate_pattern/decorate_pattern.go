package main

import "fmt"

// 抽象层
type Phone interface {
	PrintName()
}

// 抽象类 手机的装饰器，还是属于手机类型
type PhoneDecorator struct {
	Phone Phone
}

func (p *PhoneDecorator) PrintName() {}

// 实现层
type HuaweiPhone struct {
}

func (hw *HuaweiPhone) PrintName() {
	fmt.Println("这是华为手机")
}

type XiaomiPhone struct {
}

func (xm *XiaomiPhone) PrintName() {
	fmt.Println("这是小米手机")
}

type MoDecorator struct {
	PhoneDecorator
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{PhoneDecorator{phone}}
}

func (m *MoDecorator) PrintName() {
	m.PhoneDecorator.Phone.PrintName()
	m.TieMo()
}
func (m *MoDecorator) TieMo() {
	fmt.Println("手机贴膜了")
}

type KeDecorator struct {
	PhoneDecorator
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{PhoneDecorator{phone}}
}

func (k *KeDecorator) PrintName() {
	k.Phone.PrintName()
	k.DaiKe()
}
func (k *KeDecorator) DaiKe() {
	fmt.Println("手机带壳了")
}

// 业务逻辑层
func main() {
	var phone Phone
	phone = new(HuaweiPhone)
	// 给华为手机贴膜
	var moDecorator = NewMoDecorator(phone)
	moDecorator.PrintName()
	fmt.Println("-------------")
	// 给贴膜的华为手机带壳
	var keDecorator = NewKeDecorator(moDecorator)
	keDecorator.PrintName()
}
