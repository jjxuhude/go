package main

import (
	"flag"
	"fmt"
)

var skillParam = flag.String("skill", "run", "skill to perform")

func main() {
	flag.Parse()

	var skill = map[string]func(){
		"fire": func() {
			fmt.Println("chicken Fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Printf("skill not found")
	}
}
