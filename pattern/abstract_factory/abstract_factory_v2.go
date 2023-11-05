package main

import "fmt"

// 抽象工厂的作用在于不用每新建一个品类，都创建一个工厂，每个工厂生产一整套产品等级结构
// 练习：产品等级结构为
// clothes trousers shoes  衣服 裤子 鞋

// 抽象层
// 产品等级结构为Clothes，trousers，shoes
type Clothes interface {
	PutOn()
}
type Trousers interface {
	PutOn()
}
type Shoes interface {
	TakeOn()
}

// 产品族：生产完整的产品等级结构的工厂
// 一个工厂能生产出完整的产品
type AbsFactory interface {
	CreateClothes() Clothes
	CreateTrousers() Trousers
	CreateShoes() Shoes
}

// 实现层
// 不同产品族的产品等级结构全部创建完成
type ChinaClothes struct {
}

func (c *ChinaClothes) PutOn() {
	fmt.Println("穿上中国衣服")
}

type ChinaTrousers struct {
}

func (c *ChinaTrousers) PutOn() {
	fmt.Println("穿上中国裤子")
}

type ChinaShoes struct {
}

func (c *ChinaShoes) TakeOn() {
	fmt.Println("穿上中国鞋")
}

type JapanClothes struct {
}

func (j *JapanClothes) PutOn() {
	fmt.Println("穿上日本衣服")
}

type JapanTrousers struct {
}

func (j *JapanTrousers) PutOn() {
	fmt.Println("穿上日本裤子")
}

type JapanShoes struct {
}

func (j *JapanShoes) TakeOn() {
	fmt.Println("穿上日本鞋")
}

// 创建中国和日本工厂
type ChinaFactory struct {
}

func (c *ChinaFactory) CreateClothes() Clothes {
	return &ChinaClothes{}
}
func (c *ChinaFactory) CreateTrousers() Trousers {
	return &ChinaTrousers{}
}
func (c *ChinaFactory) CreateShoes() Shoes {
	return &ChinaShoes{}
}

type JapanFactory struct {
}

func (j *JapanFactory) CreateClothes() Clothes {
	return &JapanClothes{}
}
func (j *JapanFactory) CreateTrousers() Trousers {
	return &JapanTrousers{}
}
func (j *JapanFactory) CreateShoes() Shoes {
	return &JapanShoes{}
}

// 业务逻辑层
func main() {
	// 1.定义抽象中国工厂
	var chinaFac AbsFactory
	// 2.实例化中国工厂
	chinaFac = new(ChinaFactory)
	// 3.中国工厂生产中国衣服
	chinaClothes := chinaFac.CreateClothes()
	// 4.中国工厂生产中国裤子
	chinaTrousers := chinaFac.CreateTrousers()
	// 5.中国工厂生产中国鞋
	chinaShoes := chinaFac.CreateShoes()
	// 6.调用中国产品的方法
	chinaClothes.PutOn()
	chinaTrousers.PutOn()
	chinaShoes.TakeOn()
	// 日本工厂同理
	var japanFac AbsFactory
	japanFac = new(JapanFactory)
	japanClothes := japanFac.CreateClothes()
	japanTrousers := japanFac.CreateTrousers()
	japanShoes := japanFac.CreateShoes()
	japanClothes.PutOn()
	japanTrousers.PutOn()
	japanShoes.TakeOn()
}
