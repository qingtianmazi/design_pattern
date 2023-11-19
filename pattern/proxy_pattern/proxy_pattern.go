package main

import "fmt"

// 代理模式
// 定义需要购买的商品的结构
type Good struct {
	Name     string
	Producer string
}

// 抽象层
// 服务提供商的抽象化定义   向我们提供Buy(good) 方法
type Shopping interface {
	Buy(good Good)
}

// 实现层
// 韩国购物
type KoreaShopping struct {
}

func (k *KoreaShopping) Buy(good Good) {
	fmt.Println("购买了一件韩国商品：", good.Name)
}

// 美国购物
type AmericaShopping struct {
}

func (a *AmericaShopping) Buy(good Good) {
	fmt.Println("购买了一件美国商品：", good.Name)
}

// 代理类，用于封装韩国购物和美国购物对象
type ShoppingProxy struct {
	// 继承抽象结构，用于后续传入具体韩国购物对象和美国购物对象，从而实现多态功能
	shopping Shopping
}

func NewShoppingProxy(shopping Shopping) ShoppingProxy {
	return ShoppingProxy{
		shopping: shopping,
	}
}

func (s *ShoppingProxy) Buy(good Good) {
	s.shopping.Buy(good)
	s.PayTaxes(good)
}

func (s *ShoppingProxy) PayTaxes(good Good) {
	if good.Producer == "america" {
		fmt.Printf("对美国购买的商品：%v，加征%v关税\n", good.Name, "120%")
	} else if good.Producer == "korea" {
		fmt.Printf("对韩国购买的商品：%v，加征%v关税\n", good.Name, "100%")
	}
}

// 业务逻辑层
func main() {
	g1 := Good{
		Name:     "韩国化妆品",
		Producer: "korea",
	}
	g2 := Good{
		Name:     "美国生鱼片",
		Producer: "america",
	}
	k_shopping := &KoreaShopping{}
	var shopping ShoppingProxy
	shopping = NewShoppingProxy(k_shopping)
	shopping.Buy(g1)
	a_shopping := &AmericaShopping{}
	shopping = NewShoppingProxy(a_shopping)
	shopping.Buy(g2)
}
