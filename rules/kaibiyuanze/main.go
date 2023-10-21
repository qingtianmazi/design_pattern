package main

// 开闭原则

import "fmt"

/*
开闭原则：
类的改动是通过增加代码来实现的，而不是修改源代码
*/
type People interface {
	doWork()
}

type Driver struct {
}

func (d *Driver) doWork() {
	fmt.Println("汽车司机，开车")
}

type Pilot struct {
}

func (p *Pilot) doWork() {
	fmt.Println("飞行员，开飞机")
}

// 对抽象对象进行操作，用于实现多态方法
// 多态：父类指针指向子类对象，调用子类对象的方法
func PeopleDoWork(p People) {
	p.doWork()
}

func main() {
	d := Driver{}
	d.doWork()
	p := Pilot{}
	p.doWork()

	fmt.Println("---------------")

	PeopleDoWork(&Driver{})
	PeopleDoWork(&Pilot{})
}
