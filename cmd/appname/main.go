package main

import (
	"fmt"
)

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
	var X int8 = 127    // int8有符号整数的取值范围是 -128~127
	fmt.Println(X+4, X) // 对超出的范围，go语言设计了补码饶回，即：127 + 1 = -128 + 1 = -127 + 1 = -126 + 1 = -125，最后输出-125
	var Y int16 = 32767 // int16有符号整数的取值范围是：-32768~32767
	fmt.Println(Y+2, Y) // 同上，最后输出是 -32767
	var Z uint8 = 255   // uint无符号整数的取值范围是0~255，对超出的部分也会进行补码绕回
	fmt.Println(Z+1, Z) //最后输出   0
}
