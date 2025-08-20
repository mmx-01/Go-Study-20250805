package main

import "fmt"

const PI = 100

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
	//fmt.Printf("a的类型是：%d", a)
	// ================================新阶段，if else for==========================================
	//a := 102
	//if a > 1010 {
	//	fmt.Println(true)
	//} else {
	//	fmt.Println(false)
	//}
	//a := 10
	//for i := 1; i < a; i++ {
	//	fmt.Println(i)
	//}
	//i := 1
	//for {
	//	if i < a {
	//		fmt.Printf("%d ", i)
	//	}
	//	i++
	//	if i > a {
	//		break
	//	}
	//}
	// 数组
	//a := []string{"第一个元素 0", "第二个元素 1", "第三个元素 2"}
	//for index, j := range a {
	//	fmt.Println("当前索引位置：", index)
	//	fmt.Println("当前元素值", j)
	//}
	//mmap := map[int]string{
	//	1: "123",
	//	2: "456",
	//	3: "789",
	//}
	//for key, value := range mmap {
	//	fmt.Println(key, "value"+value)
	//}
	chnl := make(chan int)
	go func() {
		chnl <- 1
		chnl <- 10
		chnl <- 100
		chnl <- 1000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}
}
