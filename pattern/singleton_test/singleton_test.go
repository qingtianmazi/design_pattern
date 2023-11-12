package singleton_test

import (
	"design_pattern/pattern/singleton"
	"testing"
)

func SingleCar(sc singleton.SingletonCarInterface) {
	sc.PrintCarName()
}

func TestSingleton(t *testing.T) {
	singletonCar := singleton.GetSingleCar()
	singletonCar.PrintCarName()
	SingleCar(singletonCar)
}
