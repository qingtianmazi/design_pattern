package main

import "fmt"

// 单例模式要点：
/*
	1.某个类只能有一个实例
	2.该类必须自己创建这个实例
	3.该类必须给所有其他对象提供这个实例

	综述：保证一个类全局只能有一个实例对象，并提供一个全局访问点
*/

// 1.单例类需要是包内私有的，不能被外界访问到，否则就能实例化多个对象
type singletonCar struct {
	name string
}

// 2.访问单例对象的指针必须是私有指针，不能被外界访问到，否则外界就能修改这个指针的指向，导致单例对象丢失
var sc *singletonCar

// 饿汉模式
// 系统启动时就创建单例对象，无论后续是否需要
func init() {
	sc = newSingletonCar()
}

func newSingletonCar() *singletonCar {
	return &singletonCar{"BMW"}
}

// 3.对外提供的全局访问点，包外能够获得这个单例对象，只提供读权限，不提供写权限
func GetSingleCar() SingletonCarInterface {
	return sc
}

// GetSingleCar这个方法只能是普通全局函数，不能是单例类的成员函数
// 以下注释写法是错误的，因为无法获取到单例对象，也就无法调用获取单例对象的函数
//
//	func (sc *singletonCar) GetSingleton() *singletonCar {
//		return sc
//	}

func (sc *singletonCar) PrintCarName() {
	fmt.Println(sc.name)
}

type SingletonCarInterface interface {
	PrintCarName()
}

func main() {
	singleCar := GetSingleCar()
	singleCar.PrintCarName() // BMW
	singleCar2 := GetSingleCar()
	singleCar2.PrintCarName()            // BMW
	fmt.Println(singleCar == singleCar2) //true
}
