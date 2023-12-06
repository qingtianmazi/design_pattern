package main

import "fmt"

// 责任链模式
/*
	定义学生考学场景：
 	step1：报名考试
	step2：参加考试
	step3：通过考试
*/

// 学生
type Student struct {
	Name     string
	SignUp   bool // 报名考试
	TakeTest bool // 参加考试
	PassExam bool // 通过考试
}

type Section interface {
	Do(s *Student)           // 参与改环节
	setNext(section Section) // 设置下一个环节
}

// 学生报名考试
type SignUp struct {
	next Section
}

func (s *SignUp) Do(stu *Student) {
	if stu.SignUp {
		fmt.Println(stu.Name + "报名考试成功")
		s.next.Do(stu)
		return
	}
	stu.SignUp = true
	fmt.Println(stu.Name + "考生正在报名...")
	s.next.Do(stu)
	return
}

func (s *SignUp) setNext(section Section) {
	s.next = section
}

// 学生参加考试
type TakeTest struct {
	next Section
}

func (p *TakeTest) Do(stu *Student) {
	if stu.TakeTest {
		fmt.Println(stu.Name + "参加考试")
		p.next.Do(stu)
		return
	}
	stu.TakeTest = true
	fmt.Println(stu.Name + "考生参加考试...")
	p.next.Do(stu)
	return
}

func (p *TakeTest) setNext(section Section) {
	p.next = section
}

// 学生通过考试
// 责任链最后一个节点，不需要请求后续节点的 Do 方法了
type PassExam struct {
	next Section
}

func (p *PassExam) Do(stu *Student) {
	if stu.PassExam {
		fmt.Println(stu.Name + "通过考试成功")
		return
	}
	stu.PassExam = true
	fmt.Println(stu.Name + "通过考试...")
	return
}

func (p *PassExam) setNext(section Section) {
	p.next = section
}

func main() {
	var student = &Student{
		Name:     "张三",
		SignUp:   false,
		TakeTest: false,
		PassExam: false,
	}
	// 最后一个节点
	var pe = &PassExam{}
	// 倒数第二个节点
	var pt = &TakeTest{}
	// 倒数第三个节点
	var su = &SignUp{}
	// 把责任链的节点连接起来
	pt.setNext(pe)
	su.setNext(pt)
	// 调用责任链的首个节点的执行方法
	su.Do(student)
}
