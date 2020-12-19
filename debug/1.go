package main

import "fmt"

func main() {
	test1()
}

type stringer interface {
	string() string
}

type MyImplement struct {
}

func (m *MyImplement) String() string {
	return "hi"
}

func GetStringer() fmt.Stringer {
	var s *MyImplement = nil
	fmt.Println(s)
	//带有指针的nil值，直接返回nil
	if s == nil {
		return s
	}
	return s
}

func test1() {
	s := GetStringer()
	if s == nil {
		fmt.Println("GetStringer == nil", s)
	} else {
		fmt.Println("GetStringer != nil", s)
	}
}
