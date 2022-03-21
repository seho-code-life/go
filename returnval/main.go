package main

import "fmt"

func main() {
	// 空白符
	_, i1 := add1(10)
	fmt.Print((i1))
}

func add1(input int) (x int, y int) {
	x = input * 2
	y = input * 3
	return
}

func add(input int) int {
	return input*2 + input*3
}
