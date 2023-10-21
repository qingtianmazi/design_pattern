package main

import "fmt"

// 单一职责：每个类只对外提供一种功能
type Driver struct {
}

func (d *Driver) Drive() {
	fmt.Println("开车")
}

func (d *Driver) Fix() {
	fmt.Println("修车")
}

func main() {
	d := Driver{}
	d.Drive()
	d.Fix()
}
