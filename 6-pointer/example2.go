package main

import "fmt"

func changeValue (p *int) { // p 为指针类型变量
	*p = 10 // 解引用，从对应地址取值
}

func main() {
	var a int = 1
	changeValue(&a) // 取址
	fmt.Println("a = ", a)
}