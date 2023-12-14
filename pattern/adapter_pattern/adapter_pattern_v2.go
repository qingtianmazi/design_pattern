package main

import "fmt"

type Server interface {
	Service()
}

type ConcreteServiceA struct {
	name string
}

func (c *ConcreteServiceA) Service() {
	fmt.Println(c.name + " concreteServiceA service")
}

type ConcreteServiceB struct {
	name string
}

func (c *ConcreteServiceB) Service() {
	fmt.Println(c.name + " concreteServiceB service")
}

// 需要适配器适配service接口
type AdaptedServiceC struct {
	name string
}

func (a *AdaptedServiceC) DifferentService() {
	fmt.Println(a.name + " adaptedServiceC differentService")
}

type AdapterC struct {
	ac AdaptedServiceC
}

func (a *AdapterC) Service() {
	fmt.Println(a.ac.name + " service")
}

func main() {
	var s Server
	s = &ConcreteServiceA{"a"}
	s.Service()
	s = &ConcreteServiceB{"b"}
	s.Service()
	s = &AdapterC{AdaptedServiceC{"c"}}
	s.Service()
}
