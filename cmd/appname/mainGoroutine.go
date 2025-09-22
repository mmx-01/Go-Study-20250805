package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

// 主程序
func main() {
	//go testFor("测试并发")
	//time.Sleep(1 * time.Second)
	// 创建匿名的Goroutine
	//go func() {
	//	fmt.Println("2123231231")
	//}()
	//time.Sleep(1 * time.Second)
	// 创建通道
	//R1 := make(chan string)
	//R2 := make(chan string)
	//go protal2(R2)
	//go protal1(R1)
	//select {
	//case pot1 := <-R1:
	//	fmt.Println(pot1)
	//case pot2 := <-R2:
	//	fmt.Println(pot2)
	//}
	//ch := make(chan int)
	//go testDemo11(ch)
	//ch <- 10
	//fmt.Println(testDemo12(123))
	//newDate := time.Now()
	//fmt.Println(newDate)
	json1, _ := json.Marshal(false)
	fmt.Println(json1)

}

// 测试循环方法
func testFor(str string) {
	for i := 0; i < 10; i++ {
		fmt.Println(str)
	}
}

// 创建函数
func protal1(chanle chan string) {
	time.Sleep(2 * time.Second)
	chanle <- "protal1"
}
func protal2(chanle chan string) {
	time.Sleep(1 * time.Second)
	chanle <- "protal2"
}

// 创建一个函数
func testDemo11(ch chan int) {
	fmt.Println(1231 + <-ch)
}

// 创建函数 错误处理
func testDemo12(value int) (int, error) {
	if value < 0 {
		return 0, errors.New("value is less than 0")
	}
	return value + 2, nil
}
