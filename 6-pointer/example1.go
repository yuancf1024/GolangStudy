package main

import "fmt"

func changeValue (p int) { // p 为形参 p = 0
	p = 10
}

func main() {
	var a int = 1
	changeValue(a) // 值拷贝
	fmt.Println("a = ", a)
}