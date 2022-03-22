package main

import (
	"errors"
	"fmt"
)

func main() {
	if f, err := test(-1); err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Print(f)
	}
}
func test(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("小于0了")
	}
	return f, nil
}
