package main

import (
	"context"
	"errors"
	"fmt"
)

// 使用遍历的方式做有先后顺序的校验工作
type handler func(ctx context.Context, params map[string]interface{}) error

// 对用户token 进行校验
var checkToken = func(ctx context.Context, params map[string]interface{}) error {
	token, _ := params["token"].(string)
	if token != "123456" {
		return errors.New("token校验失败")
	}
	return nil
}

// 对用户name 进行校验
var checkUser = func(ctx context.Context, params map[string]interface{}) error {
	name, _ := params["name"].(string)
	if name != "admin" {
		return errors.New("用户校验失败")
	}
	return nil
}

// 对用户密码 进行校验
var checkPassword = func(ctx context.Context, params map[string]interface{}) error {
	pwd, _ := params["password"].(int)
	if pwd == 0 {
		return errors.New("密码校验失败")
	}
	return nil
}

func main() {
	var handlerChain = []handler{
		checkToken,
		checkUser,
		checkPassword,
	}
	ctx := context.Background()
	params := map[string]interface{}{
		"token":    "123456",
		"name":     "admin",
		"password": 0,
	}

	flag := true
	for _, handler := range handlerChain {
		err := handler(ctx, params)
		if err != nil {
			fmt.Println("校验失败" + err.Error())
			flag = false
			break
		}
	}
	if flag {
		fmt.Println("校验成功")
	}
}
