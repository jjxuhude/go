package main

import (
	"flag"
	"fmt"
)

func aaa(a int) (b int) {
	b = 2
	return
}

func main() {
	var skillParam = flag.String("skill", "", "skill to perform")
	flag.Parse()
	skill := map[string]func(){
		"one": func() {
			fmt.Println("1111")
		},
		"tow": func() {
			fmt.Println("2222")
		},
		"three": func() {
			fmt.Println("33333")
		},
	}
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("not find")
	}

}
