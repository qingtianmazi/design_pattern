package main

import "fmt"

// 单一职责：每个类只对外提供一种功能

// 司机类：专门负责开车
type Driver struct {
}

// 司机类有一个开车的方法
func (d *Driver) Drive() {
	fmt.Println("开车")
}

// 修理工类：专门负责修车
type Fixer struct {
}

// 司机类还有一个修车的方法
func (f *Fixer) Fix() {
	fmt.Println("修车")
}

func main() {
	d := Driver{}
	// 司机开车
	d.Drive() //输出：开车
	// 修理工修车
	f := Fixer{}
	f.Fix() // 输出：修车
}
