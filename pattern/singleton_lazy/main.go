package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 1.单例类需要是包内私有的，不能被外界访问到，否则就能实例化多个对象
type singletonCar struct {
	name string
}

type SingletonCarInterface interface {
	PrintName()
}

// 2.访问单例对象的指针必须是私有指针，不能被外界访问到，否则外界就能修改这个指针的指向，导致单例对象丢失
var s *singletonCar

func newSingletonCar() *singletonCar {
	return &singletonCar{
		name: "BMW",
	}
}

func (sc *singletonCar) PrintName() {
	fmt.Println(sc.name)
}

// 新增锁
var lock sync.Mutex

// 原子读操作标记位
var syncNum uint32

// 懒汉模式在获取对象的时候才会实例化对象
func GetSingleton() SingletonCarInterface {
	if atomic.LoadUint32(&syncNum) == 1 {
		return s
	}
	// 获取对象前，先加锁
	lock.Lock()
	defer lock.Unlock()
	// 不存在对象，则实例化对象
	if s == nil {
		s = newSingletonCar()
		// 对syncNum这个标记位进行复制操作
		atomic.StoreUint32(&syncNum, 1)
	}
	return s
}

func main() {
	sc := GetSingleton()
	sc.PrintName()
}
