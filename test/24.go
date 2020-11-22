package main

import "fmt"

type Invoker interface {
	Call(interface{})
}

type Struct struct {
}

func (s *Struct) Call(p interface{}) {
	fmt.Println(p)
}

func main() {
	invoker := &Struct{}
	invoker.Call("aaaa")

}
