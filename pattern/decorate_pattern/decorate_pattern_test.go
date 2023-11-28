package main

import (
	"fmt"
	"testing"
)

// 抽象 基础组件 Component
type Food interface {
	Eat() string
	Pay() int64
}

// 具体组件 ConcreteComponent
type Rice struct {
}

func (r *Rice) Eat() string {
	return "吃米饭"
}

// 支付金额
func (r *Rice) Pay() int64 {
	return 1
}

// 具体组件 ConcreteComponent
type Noodle struct {
}

func (n *Noodle) Eat() string {
	return "吃面条"
}
func (n *Noodle) Pay() int64 {
	return 2
}

// 装饰器 Decorator
type Decorator Food

// 创建一个装饰器
func NewDecorator(d Decorator) Decorator {
	return d
}

// 具体装饰器类 鸡蛋装饰器 ConcreteDecorator
type EggDecorator struct {
	Decorator
}

func NewEggDecorator(d Decorator) Decorator {
	return &EggDecorator{d}
}
func (e *EggDecorator) Eat() string {
	return e.Decorator.Eat() + " 加鸡蛋"
}
func (e *EggDecorator) Pay() int64 {
	return e.Decorator.Pay() + 3
}

// 具体装饰器类 酱汁装饰器 ConcreteDecorator
type SauceDecorator struct {
	Decorator
}

func NewSauceDecorator(d Decorator) Decorator {
	return &SauceDecorator{d}
}
func (s *SauceDecorator) Eat() string {
	return s.Decorator.Eat() + " 加酱汁"
}
func (s *SauceDecorator) Pay() int64 {
	return s.Decorator.Pay() + 4
}

func TestDecorator(t *testing.T) {
	// 基础组件 白米饭
	var rice Food
	rice = &Rice{}
	fmt.Println(rice.Eat())
	fmt.Println(rice.Pay())

	// 基础组件 面条
	var noodle Food
	noodle = &Noodle{}
	fmt.Println(noodle.Eat())
	fmt.Println(noodle.Pay())

	// 加菜了 创建一个装饰器
	var eggDecorator Decorator
	// 在白米饭上添加鸡蛋
	eggDecorator = NewEggDecorator(rice)
	fmt.Println(eggDecorator.Eat())
	fmt.Println(eggDecorator.Pay())

	var sauceDecorator Decorator
	// 继续加菜  在白米饭加鸡蛋的基础上再添加酱汁
	sauceDecorator = NewSauceDecorator(eggDecorator)
	fmt.Println(sauceDecorator.Eat())
	fmt.Println(sauceDecorator.Pay())
}
