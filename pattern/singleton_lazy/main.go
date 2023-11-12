package main

import "fmt"

type singletonCar struct {
	name string
}

var s *singletonCar

func newSingletonCar() *singletonCar {
	return &singletonCar{
		name: "BMW",
	}
}

func (sc *singletonCar) PrintName() {
	fmt.Println(sc.name)
}

func GetSingleton() *singletonCar {
	if s == nil {
		s = newSingletonCar()
	}
	return s
}

func main() {
	sc := GetSingleton()
	sc.PrintName()
}
