package main

import "fmt"

// 抽象工厂的作用在于不用每新建一个品类，都创建一个工厂
// 练习：产品等级结构为
// clothes trousers shoes  衣服 裤子 鞋

// 抽象层
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
	//创建中国工厂
	var chinaFac AbsFactory
	chinaFac = new(ChinaFactory)
	chinaClothes := chinaFac.CreateClothes()
	chinaTrousers := chinaFac.CreateTrousers()
	chinaShoes := chinaFac.CreateShoes()
	chinaClothes.PutOn()
	chinaTrousers.PutOn()
	chinaShoes.TakeOn()
	var japanFac AbsFactory
	japanFac = new(JapanFactory)
	japanClothes := japanFac.CreateClothes()
	japanTrousers := japanFac.CreateTrousers()
	japanShoes := japanFac.CreateShoes()
	japanClothes.PutOn()
	japanTrousers.PutOn()
	japanShoes.TakeOn()
}
