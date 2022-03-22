package main

import "fmt"

func main() {
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}
	// 使用range
	for i, v := range arr {
		fmt.Println(i, v)
	}

	fmt.Print(fib(10))

	// 声明一个切片使用make
	slice := make([]int, 10)
	// 也可以使用new来分配一个切片
	slice1 := new([10]int)[0:10]
	// new和make的区别就是，new会返回一个引用，make返回一个初始值，而且make只支持切片，channel，map
	fmt.Print(slice, slice1)

	// 切片巩固一下
	// 切片包含原值的引用，cap，长度3个部分，所以我们通过函数传递切片的时候，实际上也是复制了切片，但是切片还是指向原值的指针，所以这时候和直接穿引用的指针是一样的。
	// 当切片被扩充，长度不够时，此时指向的底层数组就会变了，所以更改新切片也不会影响到老切片了

}

// 用数组实现一个fib
func fib(le int) int {
	var resultArray = []int{0: 0, 1: 1}
	for i := 0; i < le; i++ {
		resultArray = append(resultArray, resultArray[len(resultArray)-1]+resultArray[len(resultArray)-2])
	}
	return resultArray[le]
}
