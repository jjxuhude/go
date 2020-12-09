package main

import (
	"bytes"
	"fmt"
)

func joinString2(strings ...string) string {
	var buf bytes.Buffer
	for _, str := range strings {
		buf.WriteString(str + " ")
	}
	return buf.String()
}

func main() {
	fmt.Println(joinString2("aaa", "bbb"))
}
