package main

import "fmt"

// const 来定义枚举类型
const (
	// 可以在const()添加一个关键字 iota，每行的iota都会累加1，第一行的iota的默认值是0
	BEIJING  = 10 * iota // iota = 0
	SHANGHAI             // iota = 1
	SHENZHEN             // iota = 2
)

const (
	a, b = iota + 1, iota + 2 // iota = 0, a = 1, b = 2
	c, d                      // c = 2, d = 3
	e, f                      // e = 3, f = 4

	g, h = iota * 2, iota * 3 // iota = 3, g = 6, h = 9
	i, k                      // i = 8, k = 12
)

func main() {
	// 常量(只读属性)
	const length int = 10

	fmt.Println("length = ", length)

	// length = 100 // 常量是不允许修改的。

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	// fmt.Println(a, b, c, d, e, f, g, h, i, k)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)

	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "k = ", k)

	// iota 只能配合const() 一起使用，iota只有在const进行累加效果
	// var a int = iota
}
