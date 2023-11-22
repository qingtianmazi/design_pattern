package main

import (
	"fmt"
	"testing"
)

// 抽象 基础类
type Food interface {
	Eat() string
	Pay() int64
}

// 具体基础类
type Rice struct {
}

func (r *Rice) Eat() string {
	return "吃米饭"
}
func (r *Rice) Pay() int64 {
	return 1
}

type Noodle struct {
}

func (n *Noodle) Eat() string {
	return "吃面条"
}
func (n *Noodle) Pay() int64 {
	return 2
}

type Decorator Food

// 装饰器
func NewDecorator(d Decorator) Decorator {
	return d
}

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
	// 基础类 白米饭
	var rice Food
	rice = &Rice{}
	fmt.Println(rice.Eat())
	fmt.Println(rice.Pay())

	// 基础类 面条
	var noodle Food
	noodle = &Noodle{}
	fmt.Println(noodle.Eat())
	fmt.Println(noodle.Pay())

	// 加菜了
	var eggDecorator Decorator
	eggDecorator = NewEggDecorator(rice)
	fmt.Println(eggDecorator.Eat())
	fmt.Println(eggDecorator.Pay())

	var sauceDecorator Decorator
	sauceDecorator = NewSauceDecorator(eggDecorator)
	fmt.Println(sauceDecorator.Eat())
	fmt.Println(sauceDecorator.Pay())
}
