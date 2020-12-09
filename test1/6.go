package main

import "fmt"

type parseError struct {
	Filename string
	Line     int
}

func (e *parseError) Error() string {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}
func newParseError(filename string, line int) error {
	return &parseError{filename, line}
}

func main() {
	var err error
	err = newParseError("main11111.go", 1)
	fmt.Println(err.Error())

	switch detail := err.(type) {
	case *parseError:
		fmt.Printf("Filename:%s Line:%d", detail.Filename, detail.Line)
	default:
		fmt.Println("其他错误")
	}

}
