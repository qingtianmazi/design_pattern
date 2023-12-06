package main

import "fmt"

// 抽象层
// 观察者
type Observer interface {
	OnObserve()
}

// 通知者
type Notifier interface {
	AddObserver(observer Observer)
	DeleteObserver(observer Observer)
	Notify()
}

// 实现层
type Student1 struct {
	status string
}

func (s1 *Student1) OnObserve() {
	fmt.Println("观察到老师来了，学生1，停止了", s1.status)
}

type Student2 struct {
	status string
}

func (s2 *Student2) OnObserve() {
	fmt.Println("观察到老师来了，学生2，停止了", s2.status)
}

type Student3 struct {
	status string
}

func (s3 *Student3) OnObserve() {
	fmt.Println("观察到老师来了，学生3，停止了", s3.status)
}

// 通知者
type Monitor struct {
	stuMap map[Observer]bool
}

// 添加观察者
func (m *Monitor) AddObserver(o Observer) {
	m.stuMap[o] = true
}

// 删除观察者
func (m *Monitor) DeleteObserver(o Observer) {
	delete(m.stuMap, o)
}

// 通知所有观察者
func (m *Monitor) Notify() {
	for o := range m.stuMap {
		o.OnObserve()
	}
}

// 业务逻辑层
func main() {
	var m = &Monitor{
		stuMap: make(map[Observer]bool),
	}
	s1 := &Student1{
		status: "打游戏",
	}
	s2 := &Student2{
		status: "刷剧",
	}
	s3 := &Student3{
		status: "化妆",
	}
	for _, v := range []Observer{s1, s2, s3} {
		m.AddObserver(v)
	}
	m.Notify()
}
