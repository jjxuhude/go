package simple

import (
	"fmt"
	"strconv"
)

func Map() {
	map1 := make(map[string]string)
	map1["name"] = "徐华德"
	map1["age"] = strconv.Itoa(15)

	map2 := map[string]string{
		"aaa": "1111",
		"bbb": "2222",
	}
	fmt.Println(map1, map2)
}
