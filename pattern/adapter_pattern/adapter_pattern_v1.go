package main

import "fmt"

// 定义客户端
type client struct {
}

func (c *client) UsingLighting(phone Phone) {
	phone.UsingLighting()
}

// 定义服务类接口
type Phone interface {
	UsingLighting()
}

type iPhone struct {
	name string
}

func (i *iPhone) UsingLighting() {
	fmt.Println(i.name + "using lighting")
}

// 安卓机无法直接匹配 lighting 接口
type Android struct {
	name string
}

func (a *Android) UsingTypeC() {
	fmt.Println(a.name + "using type-c")
}

//给安卓机添加一个适配器

type AndroidAdapter struct {
	androidPhone Android
}

func (a *AndroidAdapter) UsingLighting() {
	a.androidPhone.UsingTypeC()
}

func main() {
	var phone Phone
	phone = &iPhone{"iPhone"}
	var android Phone
	android = &AndroidAdapter{Android{"Android"}}
	cli := &client{}
	cli.UsingLighting(phone)
	cli.UsingLighting(android)
}
