package main

import "fmt"

func main() {
	// 写入defer 关键字,按照栈的方式运行，先进后出
	defer fmt.Println("main end1")

	defer fmt.Println("main end2")

	fmt.Println("main::hello go 1")
	fmt.Println("main::hello go 2")
}
