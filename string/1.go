package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "Hello,世界"
	fmt.Println(len(str))
	fmt.Print(utf8.RuneCountInString(str))
	arr1 := []byte(str)
	fmt.Println(arr1)
	arr2 := []rune(str)
	for k, v := range arr2 {
		fmt.Printf("%d:%c ", k, v)
	}

}
