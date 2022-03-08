package main

import "fmt"

// 值传递
// func swap(a int, b int) {
// 	var temp int
// 	temp = a
// 	a = b
// 	b = temp
// }

// 指针传递
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa // temp = main::a
	*pa = *pb // main::a = main::b
	*pb = temp // main::b = temp
}

func main() {
	var a int = 10
	var b int = 20

	// swap

	fmt.Println("a = ", a, "b = ", b)

	var p *int // 一级指针
	p = &a
	fmt.Println(&a)
	fmt.Println(p)

	var pp **int // 二级指针

	pp = &p
	fmt.Println(&p)
	fmt.Println(pp)
}
