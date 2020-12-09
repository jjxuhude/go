package main

import "fmt"

func accumulator(value int) func() int {
	return func() int {
		value++
		return value
	}
}

func main() {
	a := accumulator(10)
	fmt.Println(a())
	fmt.Println(a())
	fmt.Printf("%p\n", a)
	b := accumulator(20)
	fmt.Println(b())
	fmt.Println(b())
	fmt.Printf("%p", b)

}
