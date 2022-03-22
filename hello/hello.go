package main

import (
	"fmt"
	"runtime"
	"test/index"
)

const isRun = true

func main() {
	if isRun {
		fmt.Println("✅ 运行...")
	}
	fmt.Println("1")
	fmt.Println(index.Test(1, 2))
	order := &index.T{
		Id:   "1",
		Name: "1",
	}
	fmt.Println(index.AddOrder(order))
	var goos string = runtime.GOOS
	fmt.Println(goos)
	// 画一个倒立三角玩玩
	triangle(8)
	// fmt.Print(longParams("1", "sss"))
	// fmt.Println(fib(10))
	fib1 := closureFib(10)
	for i := 0; i < 10; i++ {
		fmt.Println(fib1())
	}
	fmt.Println()
}

// 参数过长
func longParams(val ...string) []string {
	return val
}

func triangle(len int) {
	// 循环len行
	// 控制行
	for i := len; i > 0; i-- {
		// 打印前面的空格
		for x := 0; x < len-i; x++ {
			fmt.Print(" ")
		}
		// 循环打印✨
		for j := 0; j < i; j++ {
			fmt.Print("✨ ")
		}
		fmt.Println("")
	}
}

func fib(l int) (res int) {
	if l <= 1 {
		res = 1
	} else {
		res = fib(l-1) + fib(l-2)
	}
	return
}

// 用closure实现fib
func closureFib(l int) (cb func() int) {
	var a int = 0
	var b int = 1
	return func() int {
		c := a
		a = b
		b = a + c
		return c
	}
}
