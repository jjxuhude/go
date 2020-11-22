package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "", "")
	flag.Parse()
	nameSlice := []byte(*name)

	fmt.Printf("%c", nameSlice)
}
