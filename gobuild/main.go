package main

import (
	"demo/gobuild/lib"
	"demo/gobuild/model"
	"fmt"
)

func main() {
	model.GetUser()
	lib.Lib1()
	lib.Lib2()
	var (
		name string
		age  int
	)
	fmt.Print("请输入名称和年龄：")
	fmt.Scan(&name, &age)
	fmt.Println("main", name, age)
}
