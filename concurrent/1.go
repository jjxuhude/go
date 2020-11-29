package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
)

func reader3(reader io.Reader) {
	scaner := bufio.NewScanner(reader)
	scaner.Split(bufio.ScanRunes)
	fmt.Printf("scaner:")
	for scaner.Scan() {
		fmt.Printf("%v", scaner.Text())
	}
}

func reader2(reader1 io.Reader) {
	buReader := bufio.NewReader(reader1)
	str, _, _ := buReader.ReadLine()
	fmt.Printf("reader2:%s", str)
}
func reader(reader1 io.Reader) {
	buReader := bufio.NewReader(reader1)
	for {
		r, size, err := buReader.ReadRune()
		_ = size
		if err == io.EOF {
			break
		}
		fmt.Printf("%c", r)

	}
}

func readN(reader1 io.Reader) {
	for {
		output1 := make([]byte, 20)
		n, err := reader1.Read(output1)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("out error:%s", err)
			}
		}
		if n > 0 {
			fmt.Printf("%s", output1[:n])
		}
	}

}

func pipe() {
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "php")

	var outBuf1 bytes.Buffer
	cmd1.Stdout = &outBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("%s", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("cmd1 wait %s", err)
	}

	cmd2.Stdin = &outBuf1
	var outBuf2 bytes.Buffer
	cmd2.Stdout = &outBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("cmd2 err:%s", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("cmd2 wait err:%s", err)
	}
	fmt.Printf("result:%s\n", outBuf2.Bytes())
}

func pipe2() {

}

func main() {
	//cmd1 := exec.Command("echo", "-n", "my linux go,爱上了房间里撒地方撒反对大师傅但是发大水方式地方上的方式地方都是发顺丰的方式对付的是发士大夫大师傅")
	//
	//reader1, err := cmd1.StdoutPipe()
	//if err != nil {
	//	fmt.Printf("Pipe Error:%s", err)
	//	return
	//}
	//if err := cmd1.Start(); err != nil {
	//	fmt.Printf("%s", err)
	//	return
	//}
	//readN(reader1)
	//reader(reader1)
	//reader2(reader1)
	//reader3(reader1)
	//fmt.Println()
	//phpipe()
	pipe2()
}
