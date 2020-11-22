package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func joinString(arr ...interface{}) string {
	var strBuf bytes.Buffer
	var str string
	for _, arg := range arr {
		switch arg.(type) {
		case string:
			str = arg.(string)
		case int:
			str = strconv.Itoa(arg.(int))
		}
		strBuf.WriteString(str + " ")

	}
	return strBuf.String()
}

func joinString1(arr ...interface{}) string {
	var strBuf bytes.Buffer
	var str string
	for _, arg := range arr {
		str = fmt.Sprintf("%v", arg)
		strBuf.WriteString(str + " ")

	}
	return strBuf.String()
}

func main() {
	str := joinString1("aaa", "bbb", "ccc", 1, 2, 3)
	fmt.Println(str)
}
