package main

import (
	"context"
	"errors"
	"fmt"
)

type Student1 struct {
	Name     string
	SignUp   bool // 是否考试报名
	TakeTest bool // 是否参加考试
	PassExam bool // 是否通过考试
}

type handler1 func(ctx context.Context, stu *Student1) error

// 报名考试
func SignUp1(ctx context.Context, stu *Student1) error {
	if stu.SignUp {
		fmt.Println(stu.Name + "已经报名了考试")
		return errors.New(stu.Name + "已经报名了考试了")
	}
	stu.SignUp = true
	fmt.Println(stu.Name + "正在进行考试报名")
	return nil
}

// 参加考试
func TakeTest1(ctx context.Context, stu *Student1) error {
	if stu.TakeTest {
		fmt.Println(stu.Name + "已经参加了考试")
		return errors.New(stu.Name + "已经参加过考试了")
	}
	stu.TakeTest = true
	fmt.Println(stu.Name + "正在参加考试")
	return nil
}

// 通过考试
func PassExam1(ctx context.Context, stu *Student1) error {
	if stu.PassExam {
		fmt.Println(stu.Name + "已经通过了考试")
		return errors.New(stu.Name + "已经通过了考试")
	}
	stu.PassExam = true
	fmt.Println(stu.Name + "通过了考试")
	return nil
}

func main() {
	var handlers = []handler1{SignUp1, TakeTest1, PassExam1}
	var student = &Student1{
		Name: "张三",
	}
	for _, handler := range handlers {
		if err := handler(context.Background(), student); err != nil {
			fmt.Println(err)
		}
	}
}
