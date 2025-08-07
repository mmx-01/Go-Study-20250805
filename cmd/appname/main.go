package main

import "fmt"

func main() {
	//fmt.Println("测试")
	//fmt.Println(pkg1.Demo1())
	//var name = "limengjie"
	//fmt.Println(name)
	//var _name = "limengjies"
	//fmt.Println(_name)
	//const (
	//	A = iota
	//	B = iota
	//)
	//fmt.Printf("A: %s\n", A)
	//fmt.Println("B", B)
	//var X int8 = 127    // int8有符号整数的取值范围是 -128~127
	//fmt.Println(X+4, X) // 对超出的范围，go语言设计了补码饶回，即：127 + 1 = -128 + 1 = -127 + 1 = -126 + 1 = -125，最后输出-125
	//var Y int16 = 32767 // int16有符号整数的取值范围是：-32768~32767
	//fmt.Println(Y+2, Y) // 同上，最后输出是 -32767
	//var Z uint8 = 255   // uint无符号整数的取值范围是0~255，对超出的部分也会进行补码绕回
	//fmt.Println(Z+1, Z) //最后输出   0
	//a := 20.45
	//b := 100.32
	//fmt.Println(a + b)
	//z := 3.0 + 4.0i
	//module := math.Hypot(real(z), imag(z))
	//fmt.Println("module: ", module)
	//fmt.Println(real(z), imag(z))
	//a := 123
	//b := "123"
	//c := "456"
	//result := a == (b + c)
	//fmt.Println(result)
	//fmt.Println(len(a))
	//var _Customer = "123"
	//var _A = 123
	//var _B = 123.123
	//fmt.Printf("%T\n", _Customer)
	//fmt.Printf("%T\n", _A)
	//fmt.Printf("%T\n", _B)
	//var _A, _b, _C = 1, "123", 11.11
	////fmt.Println(_A, _b, _C)
	//fmt.Printf("_A的类型为：%T\n", _A)
	//fmt.Println("_A的值为：", _A)
	//fmt.Printf("_b的类型为：%T\n", _b)
	//fmt.Println("_b的值为：", _b)
	//fmt.Printf("_C的类型为：%T\n", _C)
	//fmt.Println("_C的值为：", _C)
	// 使用短变量声明
	Customer := "limengjie"
	age := 27
	gongzi := 23000.91
	fmt.Printf("Customer的类型为：%T\n", Customer)
	fmt.Println("Customer的值为：", Customer)
	fmt.Printf("age的类型为：%T\n", age)
	fmt.Println("age的值为：", age)
	fmt.Printf("gongzi的类型为：%T\n", gongzi)
	fmt.Println("gongzi的值为", gongzi)
	//fmt.Printf("a的类型是：%d", a)
}
