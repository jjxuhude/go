package main

import (
	"fmt"
	"os"
)

func fileSize(filename string) int64 {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fileStat, err := f.Stat()
	if err != nil {
		panic(err)
	}

	return fileStat.Size()

}

func main() {
	var filename = os.Args[1]

	fmt.Println(fileSize(filename))
}
