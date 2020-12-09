package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var data = map[string]string{"one": "1", "two": "2"}

func readValue(keyName string) string {
	var val string
	lock.Lock()
	defer lock.Unlock()

	val = data[keyName]
	return val
}

func main() {
	fmt.Println(readValue("one"))
}
