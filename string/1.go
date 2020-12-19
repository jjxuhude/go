package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//str1();
	//str2();
	str3()

}

func str3() {
	str := "hello world !"
	println("包含：", strings.Contains(str, "world"))
	println("数量：", strings.Count(str, "o"))
	println("字段：", strings.Fields(str))
	println("前缀：", strings.HasPrefix(str, "hello"))
	println("后缀：", strings.HasSuffix(str, "!"))
	println("位置：", strings.Index(str, "world"))
	println("join：", strings.Join([]string{"hello", "world"}, "-"))
}

func str2() {
	str := "Hello,世界"
	strArr := []byte(str)
	//strArr2 := []rune(str)
	string := string(strArr)
	fmt.Println(str, strArr, string)
}

func str1() {
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
