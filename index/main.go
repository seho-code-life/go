package index

import "fmt"

type IZ int

type T struct {
	Id   string
	Name string
}

func main() {
	fmt.Println("hello, world")
}

func Test(a int, b int) int {
	fmt.Println("测试")
	return a + b
}

func AddOrder(p *T) bool {
	return true
}
