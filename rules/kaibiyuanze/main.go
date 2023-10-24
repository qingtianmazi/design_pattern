package main

import "fmt"

type People struct {
}

func (p *People) Drive() {
	fmt.Println("司机开车")
}

func (p *People) Fly() {
	fmt.Println("飞行员开飞机")
}

func (p *People) Ship() {
	fmt.Println("船员开船")
}

func main() {
	var p People
	p.Drive() //司机开车
	p.Fly()   //飞行员开飞机
}

//type People interface {
//	doWork()
//}
//
//type Driver struct {
//}
//
//func (d *Driver) doWork() {
//	fmt.Println("司机开车")
//}
//
//type Pilot struct {
//}
//
//func (p *Pilot) doWork() {
//	fmt.Println("飞行员开飞机")
//}
//
//// 对抽象对象进行操作，用于实现多态方法
//// 多态：父类指针指向子类对象，调用子类对象的方法
//func PeopleDoWork(p People) {
//	p.doWork()
//}
//
//func main() {
//	d := Driver{}
//	d.doWork() // 司机开车
//	p := Pilot{}
//	p.doWork() // 飞行员开飞机
//
//	fmt.Println("---------------")
//
//	PeopleDoWork(&Driver{}) // 司机开车
//	PeopleDoWork(&Pilot{})  // 飞行员开飞机
//}
