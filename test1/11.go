package main

import "fmt"

type BaseColor struct {
	R, G, B float32
}

type Color struct {
	BaseColor
	Alpha float32
}

func main() {
	//var c = Color{Base: BaseColor{11,22,33},Alpha: 1}
	var c Color
	c.R = 1
	c.G = 2
	c.B = 3
	c.Alpha = 4
	fmt.Printf("%v", c)
}
