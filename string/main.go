package main

import (
	"fmt"
	"strings"
)

const test = "hello world"

func main() {
	fmt.Println(strings.HasPrefix(test, "hl"))
}
