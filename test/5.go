package main

import (
	"errors"
	"fmt"
)

var err = errors.New("不能为0")

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, err
	}
	return a / b, nil
}

func main() {
	fmt.Println(div(2, 0))
}
