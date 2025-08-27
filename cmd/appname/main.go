package main

import (
	"fmt"
	"strconv"
	"strings"
)

const PI = 100

// 结构体
type person struct {
	name string
	age  int
}

// 创建两个结构体，学生和老师
type student struct {
	name   string
	age    int
	branch string
}

type teacher struct {
	name     string
	language string
	per      person
}

// 非结构体
type p1 int

// 匿名字段
type mal struct {
	string
	int
	bool
}

// 创建一个函数，用作结构体的参数
type testFunc func(int, int) int

// 创建结构体
type phone struct {
	name string
	age  int
	fun  testFunc
}

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
	//chnl := make(chan int)
	//go func() {
	//	chnl <- 1
	//	chnl <- 10
	//	chnl <- 100
	//	chnl <- 1000
	//	close(chnl)
	//}()
	//for i := range chnl {
	//	fmt.Println(i)
	//}
	// 循环控制语句  break goto  continue
	// break
	//a := 10
	//for i := 0; i <= a; i++ {
	//	if i > a-5 {
	//		break
	//	}
	//	fmt.Println(i)
	//}
	// continue
	//var a = 10
	//for i := 0; i <= a; i++ {
	//	if i == a-5 {
	//		continue
	//	}
	//	fmt.Println(i)
	//
	//}
	// switch   带有可选参数的方法
	//switch a := 3; a {
	//case 1:
	//	fmt.Println("1")
	//case 2:
	//	fmt.Println(2)
	//case 3:
	//	fmt.Println(3)
	//
	//}
	// switch 不带可选参数
	//var value = 10
	//switch {
	//case value == 1:
	//	fmt.Println(1)
	//case value == 2:
	//	fmt.Println(2)
	//case value == 3:
	//	fmt.Println(3)
	//case value != 4:
	//	value += 1
	//	if true {
	//		fmt.Println(value)
	//	}
	//case value == 11:
	//	fmt.Println(value)
	//}
	//fmt.Println(value)
	/**
	* 死锁：当您试图从通道读取或写入数据但通道没有值时。
	因此，它阻塞goroutine的当前执行，并将控制传递给其他goroutine，
	但是如果没有其他goroutine可用或其他goroutine睡眠，由于这种情况，
	程序将崩溃。这种现象称为死锁。如下面的实例所示：
	*/
	//c := make(chan int) // 创建通道，这样的创建方式是可以直接使用的，例如以下例子
	//select {
	//case <-c: // 若在select语句中，没有设置默认的case语句，出现死锁的情况就会报错
	//default:
	//	fmt.Println("执行默认的select case语句")
	//}
	//// 创建一个直接可用的通道
	//cc := make(chan string)
	//
	//go func() {
	//	cc <- "给cc通道传递了一个值"
	//}()
	//
	//value := <-cc
	//fmt.Println(value)
	//var value =
	//fmt.Printf("矩形的面积为：%d", testDemo1(123, 456))
	//fmt.Println(testDemo2(`123`, `456`))
	// 匿名函数
	//func() {
	//	fmt.Println("hello world")
	//}()
	//func(parameter_list) return_type {
	//	//代码
	//	//如果给定return_type，则使用return语句
	//	//如果未提供return_type，则不
	//	//使用return语句
	//	return
	//}()
	// 匿名函数可以直接分配给变量，如下所示

	//result := func() {
	//	fmt.Println("Hello World,这是将一个匿名函数分配给一个变量")
	//}
	//result()

	//fmt.Println(testDemo2("123", "456"))
	// 在匿名函数中，传递参数
	//resule := func(ele string) string {
	//	return ele
	//}("测试匿名函数传递参数")
	//fmt.Println(resule)
	// 将匿名函数当做参数传递给匿名函数
	//result := func(q, p string) string {
	//	return q + "_" + p
	//}
	//Lee(result)
	//fmt.Println(result)
	// 将匿名函数作为return 返回
	//result := GFG()
	//fmt.Println(result("Frist params 'i'", "Second params 'j'"))

	//Z := 120
	//
	//fmt.Printf("Z值函数调用之前：%d\n", Z)
	//
	//Demo4(Z)
	//
	//fmt.Printf("Z值函数调用之后：%d", Z)
	//var x int = 10
	//var y int = 20
	//fmt.Println("交换前的x的值：" + strconv.Itoa(x) + "交换前y的值：" + strconv.Itoa(y))
	//swap(x, y)
	//fmt.Println("交换后的x的值：" + strconv.Itoa(x) + "交换后y的值：" + strconv.Itoa(y))
	//fmt.Println(y)
	//swapData(&x, &y) // 交换成功，直接将引用传递
	//fmt.Println(y)
	//var a, b, c = returnDemo1(2, 4) // 函数接收两个参数，返回三个参数，将三个参数分别赋值到三个变量
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//var a, _ = mul_div(1, 2)
	//fmt.Println(a)
	//mul(10, 10)       // 1
	//defer mul(20, 20) // 2
	//show()            // 3
	// 上述调用函数，执行顺序为 1、3、2 因为使用了defer关键字，延迟函数调用
	// 直到延迟函数附近所有的函数执行完毕后，才会执行延迟函数
	//res := person{
	//	name: "123",
	//	age:  123,
	//}
	//showPerson(res)
	//value1 := p1(123)
	//value2 := p1(2)
	//res := value2.showp1(value1)
	//fmt.Println(res)
	//val1 := teacher{"小红", "英语、语文、这个老师是班主任"}
	//val2 := student{"小明", 21, "计算机科学与技术"}
	//val1.showDemo()
	//val2.showDemo()
	//result := person{"李梦杰", 123}
	////result.name = "xzw"
	//result.askPerson()
	//tmp := person{"leemengjie", 13}
	//tmp1 := &tmp
	//fmt.Println((*tmp1).name) // 显示指针调用
	//fmt.Println(tmp1.name)    // go中的语法糖
	//fmt.Println(tmp.name)     // 普通的方法调用
	// 以上的区别为：使用指针类型：从内存中找到person变量，然后解引用，从解引用后的person变量中找到name 显示调用和语法糖在内存中执行是一样的，只不过语法简洁
	// 直接调用方法：找到内存中person变量，直接从变量中找到name的值
	// 结构体比较 ==
	//per1 := person{"li", 123}
	//per2 := person{"li", 123}
	//per3 := person{"li", 1213}
	//fmt.Println(per1 == per2)
	//fmt.Println(per2 == per3)
	// 结构体比较  使用DeepEqual()方法
	//per1 := person{"lee", 1}
	//per2 := person{"lee", 2}
	//fmt.Println(reflect.DeepEqual(per1, per2))
	//result := teacher{per: person{"123", 11}}
	//result.askDemo2()
	//Element := struct {
	//	name string
	//	age  int
	//}{
	//	"姓名",
	//	123,
	//}
	//fmt.Println(Element)
	//result := mal{"姓名", 133, false}
	//result.showMal()
	//testFunc(1111, 1111)
	//result := phone{"测试函数类型属性", 222, func(y int, x int) int {
	//	return y + x
	//}}
	////fmt.Println("函数方法结果：", result.fun(result.age, result.age))
	////fmt.Println("age：", result.age)
	////fmt.Println("name：", result.name)
	//result.Demo6()
	//result := Demo7()
	//for i := 1; i <= len(result); i++ {
	//	fmt.Println(result[i-1])
	//}
	//fmt.Println(len(result))
	//arr := [4]string{"231231", "2123", "443", "121"}
	//fmt.Println(arr)
	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}
	// 创建多维数组  创建多维数组的时候，第一个[4]表示可以创建几个数组，第二个[3]表示数组的长度
	//arr := [4][3]int{
	//	{1, 2, 3},
	//	{5, 6, 7},
	//	{7, 8, 9},
	//	{3, 4, 2},
	//}
	//fmt.Println(arr)
	//for i, v := range arr {
	//	fmt.Println(i, v)
	//}
	// 数组复制
	//arr := [...]string{"a", "b", "c", "d", "e", "f", "g"}
	//fmt.Println("原始数组", arr)
	//newArr := arr
	//newArr[0] = "0000"
	//fmt.Println("已更改的复制的数组", newArr)
	// 通过引用复制数组
	//arr := [...]int{1, 2, 3, 4, 5}
	//fmt.Println("原始数组：", arr)
	//newArr := &arr
	//arr[1] = 10
	//fmt.Println("通过引用复制的数组：", newArr)
	////for i, v := range newArr {
	////	fmt.Println(i, v)
	////}
	//fmt.Println(*newArr)
	// 调用数组为参数的函数
	//var arr = [...]string{"a", "b", "c", "d", "e", "f", "g"}
	//result := Demo8(arr)
	//fmt.Println(result)
	// 创建切片
	// 创建数组
	var arr = [5]int{1, 2, 3, 4, 5}
	myslice := arr[1:3] // 解释：通过数组创建一个切片，1表示从索引1开始，3表示切片大小，由于是从数组的索引1开始取的
	fmt.Println(myslice)
}

// 若函数在参数后写了return_type类型，那么此函数必须要使用return来返回结果
func testDemo1(wdith int, height int) int {

	return wdith * height

	//fmt.Println(str)
}

// 可变参数  若函数的参数写成可变参数，类型就变成数组类型了
func testDemo2(element ...string) []string {
	// 创建字符串切片的方式
	// fruits := []string{"apple", "banana", "cherry"}
	return element
}

// 创建一个名为Lee的函数，此函数的参数是一个匿名函数，i 是这个匿名函数的指针，表示指向这个函数，也可以说是函数的引用/变量
// 这个匿名函数接收两个string类型的参数，q和p，函数Lee没有返参类型

func Lee(i func(q, p string) string) {
	fmt.Println(i("123", "456"))
}

//	func Lee(i func(q, p string) string) {
//		│      │  │   │  │     │      │
//		│      │  │   │  │     │      └── 内层函数i的返回类型：string
//		│      │  │   │  │     └───────── 参数p的类型：string
//		│      │  │   │  └─────────────── 参数q的类型：string
//		│      │  │   └────────────────── 内层函数i的参数列表
//		│      │  └────────────────────── 内层函数i的声明
//		│      └───────────────────────── 参数i的名称
//		└──────────────────────────────── 外层函数Lee的声明
//
// 返回匿名函数，go中不仅可以将一个匿名函数作为参数，并且还要将一个匿名函数使用return返回
func GFG() func(i, j string) string {
	value := func(i, j string) string {
		return i + j + "测试返回一个匿名函数"
	}
	return value
}

// main 函数
func Demo3() int {
	// 切片创建
	//s := []int{6, 1, 3, 4, 5, 2}
	//// 对切片进行排序
	//sort.Ints(s)
	//return s
	// 查询索引值
	str := "mengjie"
	return strings.Index(str, "m") // 此方法会返回第一个匹配到的字符，并返回索引值（从0开始）
}

// init函数 执行顺序是在main之前，用途：可以在main函数未执行之前，初始化一些全局变量
//func init() {
//	fmt.Println("b")
//}
//
//func init() {
//	fmt.Println("a")
//}

//
/**
函数参数  实际参数不能通过函数调用的方式通过形参改变其值
当函数传递过来一个参数的时候，传递的是参数的副本，而不是参数本身，所有即便
传递过来参数也是不能改变的
*/
// 修改函数值，修改失败 原因如上

func Demo4(Z int) {
	Z = 10
}

// 交换函数值
func swap(x, y int) string {
	// 设置一个临时变量   x = 10 y = 20
	var tmp int
	tmp = x
	x = y
	y = tmp
	return "交换后的x的值为：" + strconv.Itoa(x) + ", 修改后的y的值：" + strconv.Itoa(y)
}

// 引用调用
func swapData(x, y *int) int {
	var tmp int
	tmp = *x
	x = y
	*y = tmp
	return tmp
}

// return返回多个值
func returnDemo(i, y int) (int, int, int) {
	return i + y, i * y, i - y
}

// 给return返回值命名
func returnDemo1(i, y int) (name1, name2, name3 int) {
	name1 = i * y
	name2 = i - y
	name3 = i * i
	return name1, name2, name3
}

// 空白标识符
func mul_div(i, j int) (name1, name2 int) {
	name1 = i * j
	name2 = j - i
	return name1, name2
}

// defer关键字  延迟函数调用
func mul(i, j int) int {
	reslut := i * j
	fmt.Println(reslut)
	return 0
}
func show() {
	fmt.Println("测试延迟函数调用")
}

// 方法 结构体
func showPerson(a person) {
	fmt.Println(a.name)
	fmt.Println(a.age)
}

// 方法 非结构体
func (i p1) showp1(d2 p1) p1 {
	return i * d2
}

// 创建两个相同的方法，但是接收参数类型不同
func (stu student) showDemo() {
	fmt.Println("学生姓名：" + stu.name)
	fmt.Println("学生年龄：" + strconv.Itoa(stu.age))
	fmt.Println("所学学科：" + stu.branch)
}
func (tea teacher) showDemo() {
	fmt.Println("老师姓名：" + tea.name)
	fmt.Println("所授科目：" + tea.language)
}

// 访问结构体
//func (per person) askPerson() {
//	fmt.Println(per.name)
//	fmt.Println(per.age)
//}
//
//func (tea teacher) askDemo2() {
//	fmt.Println("打印嵌套结构体中person的name：" + tea.per.name)
//}
//
//// 匿名字段
//func (m mal) showMal() {
//	fmt.Println("字符串类型", m.string)
//	fmt.Println("数字类型", m.int)
//	fmt.Println("布尔类型", m.bool)
//}
//
//func limengjie(a, b, c string) string {
//	return a + b + c
//}
//
//// 测试结构体中嵌套函数类型参数
//func (ph phone) Demo6() {
//	fmt.Println(ph.name)
//	fmt.Println(ph.age)
//	fmt.Println(ph.fun(ph.age, ph.age))
//	//fmt.Println(ph.fun())
//}
//
//// 创建数组
//
//func Demo7() [2]string {
//	var studntArray [2]string
//	studntArray[0] = "1321"
//	studntArray[1] = "1322"
//	return studntArray
//}

// 将数组当做函数参数传递
func Demo8(a [7]string) [7]string {
	return a
}
