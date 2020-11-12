package simple

import (
	"fmt"
	"reflect"
)

func Ptr(a, b *int) {
	*a, *b = *b, *a
}

func Interface(args ...interface{}) {
	for _, arg := range args {
		argType := reflect.TypeOf(arg).Kind()
		fmt.Println(argType)
	}
}
