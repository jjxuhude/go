package simple

import (
	"fmt"
)

func Substr(str string) string {
	txt := "abc好好学习"
	for k, v := range []rune(txt) {
		fmt.Printf("k:%d,v:%c \n", k, v) //直接打印v的话是ascii码对照表
	}
	return ""
}
