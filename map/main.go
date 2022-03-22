package main

import "fmt"

func main() {
	var maplit map[string]int
	// 不要使用new创造map，要使用make，因为map是引用类型，如果使用new创造了map，那么返回就是一个空指针，在上面修改会直接报错
	maplit = map[string]int{"one": 1}
	fmt.Print(maplit["one"])
}
