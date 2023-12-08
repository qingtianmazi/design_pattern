package main

import "fmt"

// 抽象层
// 发布者接口
type Publisher interface {
	Follow(s Subscriber)
	UnFollow(s Subscriber)
	Notify()
}

// 订阅者接口
type Subscriber interface {
	Update()
	Name() string
}

// 实现层
// 具体发布者
type ConcretePublisher struct {
	subscribers []Subscriber
}

func (c *ConcretePublisher) Follow(s Subscriber) {
	c.subscribers = append(c.subscribers, s)
}

func (c *ConcretePublisher) UnFollow(s Subscriber) {
	c.subscribers = remove(c.subscribers, s)
}

func (c *ConcretePublisher) Notify() {
	for _, s := range c.subscribers {
		s.Update()
	}
}

func remove(s []Subscriber, s1 Subscriber) []Subscriber {
	for i, concreteSubscriber1 := range s {
		if concreteSubscriber1.Name() == s1.Name() {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

// 具体订阅者
type ConcreteSubscriber struct {
	name string
}

func (c *ConcreteSubscriber) Update() {
	fmt.Println(c.name + "收到订阅通知")
}

func (c *ConcreteSubscriber) Name() string {
	return c.name
}

// 业务逻辑层
func main() {
	var publisher Publisher
	publisher = &ConcretePublisher{}
	var subscriber1 Subscriber
	var subscriber2 Subscriber
	subscriber1 = &ConcreteSubscriber{"张三"}
	subscriber2 = &ConcreteSubscriber{"李四"}
	publisher.Follow(subscriber1)
	publisher.Follow(subscriber2)
	publisher.Notify()
}
