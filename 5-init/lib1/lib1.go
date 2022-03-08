package lib1

import "fmt"

// 当前Lib1包提供的API
func Lib1Test() {
	fmt.Println("lib1Test()...")
}

func init() {
	fmt.Println("lib1. init() ...")
}
