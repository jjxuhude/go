package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func reader() {
	str := "我的名字叫徐华德，jack.xu1"
	bReader := bufio.NewReader(strings.NewReader(str))
	for {
		c, size, err := bReader.ReadRune()
		fmt.Printf("%c", c)
		_ = size
		if err == io.EOF {
			break
		}
	}
	fmt.Println()

}

func scaner() {
	input := "hello 中文 world"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		fmt.Printf("%v", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

}

func writer() {
	buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(buf)
	w.WriteString("你好呀中国\n")
	w.Flush() // 不要忘记刷新
	fmt.Printf("\n%s", buf.String())
}

func main() {

	reader()
	scaner()
	writer()

}
