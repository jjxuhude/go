package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func strReader() {
	reader := strings.NewReader("abcdefghij")
	fmt.Println("len:", reader.Len())
	fmt.Println("size:", reader.Size())

	var buf1 = make([]byte, 5)
	reader.Read(buf1)
	fmt.Println(string(buf1))
	fmt.Println("len:", reader.Len())
	fmt.Println("size:", reader.Size())
}

func bytesReader() {
	bytes1 := []byte("我是中国工人")
	reader := bytes.NewReader(bytes1)
	fmt.Println("len:", reader.Len())
	fmt.Println("size", reader.Size())

	var buf1 = make([]byte, 5)
	reader.Read(buf1)
	fmt.Println(string(buf1))
	fmt.Println("len:", reader.Len())
	fmt.Println("size:", reader.Size())

}

func bytesBuffer() {
	var str = "我是aabc,中国工人"
	var bArr = []byte(str)
	var reader = bytes.NewBuffer(bArr)
	fmt.Println("len:", reader.Len())
	fmt.Println("size:", reader.Cap())

	var buf1 = make([]byte, 7)
	reader.Read(buf1)
	fmt.Println(string(buf1))
	fmt.Println("len:", reader.Len())
	fmt.Println("cap:", reader.Cap())
}

func bufioReader() {
	var strReader = strings.NewReader("abcdefghijklmn")
	var bufioReader = bufio.NewReaderSize(strReader, 16)
	var buf = make([]byte, 10)
	n, _ := bufioReader.Read(buf)
	fmt.Printf("未读：%d,已读：%s", bufioReader.Buffered(), buf[:n])
}

func bufioReader1() {
	var strReader = strings.NewReader("你好呀：abcdefghijklmn")
	var bufioReader = bufio.NewReaderSize(strReader, 16)
	for {
		s, _, err := bufioReader.ReadRune()
		fmt.Printf("%c", s)
		if err == io.EOF {
			break
		}

	}
}

func bufioReader2() {
	f, _ := os.Open("../fmt/1.go")
	defer f.Close()
	var bufioReader = bufio.NewReader(f)
	for {
		s, _, err := bufioReader.ReadRune()
		fmt.Printf("%c", s)
		if err == io.EOF {
			break
		}

	}
}

func main() {
	//strReader()
	//bytesReader()
	//bytesBuffer()
	//bufioReader()
	//bufioReader1()
	bufioReader2()

}
