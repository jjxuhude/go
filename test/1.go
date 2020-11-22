package main

import (
	"fmt"
	"runtime"
)

var m = map[int]string{1: "A", 2: "B", 3: "C"}
var info string

func init() {
	info = fmt.Sprintf("OS:%s,Arch:%s", runtime.GOOS, runtime.GOARCH)
}

func PlayGen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}

}

func main() {
	user := PlayGen("jack.xu1")
	name, hp := user()
	fmt.Println(name, hp)
}
