package main

import (
	"fmt"
)

// 第二个main，测试函数
func main() {

	// 调用接口
	//var val fristInterface  定义一个接口的变量
	//val = study{1, 2}  将结构体赋值给接口变量，前提是，在此接口中的实现方法的参数是该结构体
	//fmt.Println("两数相加：", val.Volume())
	//fmt.Println("两数相乘：", val.Tarea())
	//var a interface{} = "123"
	//mysel(a)
	// 给接口定义变量
	//var in testDemoInterface
	//in = testDemo{111, 222}
	//fmt.Println(in.add())
	var fileInterface ReadWriteCloser
	fileInterface = FileString{"这是读接口", "这是写接口", "这是关闭接口"}
	fmt.Println(fileInterface.Read())
	fmt.Println(fileInterface.Write())
	fmt.Println(fileInterface.Closer())
}

// 定义一个接口
type fristInterface interface {
	// 方法
	Tarea() float64
	Volume() float64
}

// 创建一个结构体
type study struct {
	name float64
	age  float64
}

// 实现接口中的方法
func (stu study) Tarea() float64 {
	return stu.name * stu.age
}
func (stu study) Volume() float64 {
	return stu.age + stu.name
}

// 定义一个方法，接收一个接口入参
func mysel(a interface{}) {
	value, ok := a.(string)
	fmt.Println(value, ok)
}

// 定义一个接口
type testDemoInterface interface {
	add() int
}

// 定义结构体
type testDemo struct {
	one int
	two int
}

// 实现方法
func (td testDemo) add() int {
	return td.one + td.two
}

// 定义一个读
type Read interface {
	Read() string
}

// 写
type Write interface {
	Write() string
}

// 关闭
type Closer interface {
	Closer() string
}

// 接口嵌套 定义一个大接口，来实现读写关闭功能
type ReadWriteCloser interface {
	Read
	Write
	Closer
}

// 实现方法
func (f FileString) Read() string {
	return f.read
}
func (f FileString) Write() string {
	return f.write
}
func (f FileString) Closer() string {
	return f.close
}

// 结构体
type FileString struct {
	read  string
	write string
	close string
}
