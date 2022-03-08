package main

import (
	// "lib1" // 错误导包
	// "lib2"

	_ "github.com/yuancf1024/GolangStudy/5-init/lib1"
	mylib2"github.com/yuancf1024/GolangStudy/5-init/lib2"
	// ."github.com/yuancf1024/GolangStudy/5-init/lib2"
)

func main() {
	// lib1.Lib1Test()
	// lib2.Lib2Test()
	// mylib2.Lib2Test()
	mylib2.Lib2Test()
}
