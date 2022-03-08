// 知识点二：defer和return谁先 谁后

package main

import "fmt"

func deferFunc() int {
	fmt.Println("defer fun called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called.")
	return 0
}

func returnAndDefer() int {
	defer deferFunc() // 当前函数全部执行完后，才出栈执行
	return returnFunc()
}

func main() {
	returnAndDefer()
}
