package main

import "fmt"

func main() {
	fmt.Print(callback(6, Add))
}

func Add(a int, b int) int {
	return a * b
}

func callback(i int, callback func(a int, b int) int) int {
	return callback(i, 10)
}
