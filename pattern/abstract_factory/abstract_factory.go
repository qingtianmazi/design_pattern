package main

import "fmt"

// 产品族：组成一个整体的所有部分，如：电脑是一个整体产品，cpu，内存，硬盘是部分，电脑由这三个部分组成（假设电脑只有这三部分）
// 产品等级结构：每个部分都有很多生产厂商，如：cpu有英特尔，ARM两个生产厂商，这两个生产厂商生产的cpu就组成一个产品等级结构

// 抽象层
// 一个产品族就有这三种产品
type CPU interface {
	Show()
}
type Mem interface {
	Show()
}
type Disk interface {
	Show()
}

// 抽象工厂，生产产品族
type AbstractFac interface {
	CreateCPU() CPU
	CreateMem() Mem
	CreateDisk() Disk
}

// 实现层
// 实现具体哪个厂商的零部件
type IntelCPU struct {
}

func (i *IntelCPU) Show() {
	fmt.Println("intelCPU")
}

type IntelMem struct {
}

func (i *IntelMem) Show() {
	fmt.Println("intelMem")
}

type IntelDisk struct {
}

func (i *IntelDisk) Show() {
	fmt.Println("intelDisk")
}

type KingstonCPU struct {
}

func (k *KingstonCPU) Show() {
	fmt.Println("kingstoncpu")
}

type KingstonMem struct {
}

func (k *KingstonMem) Show() {
	fmt.Println("kingstonmem")
}

type KingstonDisk struct {
}

func (k *KingstonDisk) Show() {
	fmt.Println("kingstondisk")
}

// 具体生产工厂
type IntelFac struct {
}

func (i *IntelFac) CreateCPU() CPU {
	var cpu CPU
	cpu = new(IntelCPU)
	return cpu
}
func (i *IntelFac) CreateMem() Mem {
	var mem Mem
	mem = new(IntelMem)
	return mem
}
func (i *IntelFac) CreateDisk() Disk {
	var disk Disk
	disk = new(IntelDisk)
	return disk
}

// Kingston 同理
type KingstonFac struct {
}

// 业务逻辑层
func main() {
	var fac AbstractFac
	var cpu CPU
	fac = new(IntelFac)
	cpu = fac.CreateCPU()
	cpu.Show()
	var mem Mem
	mem = fac.CreateMem()
	mem.Show()
	var disk Disk
	disk = fac.CreateDisk()
	disk.Show()
}
