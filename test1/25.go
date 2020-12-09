package main

import "fmt"

//函数体实现接口
type Invoker1 interface {
	Call(interface{})
}

type FunCaller func(interface{})

func (f FunCaller) Call(p interface{}) {
	f(p)
}

func main() {
	var invoker1 Invoker1
	invoker1 = FunCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	invoker1.Call("hello")
}
