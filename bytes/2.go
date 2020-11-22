package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

func main() {
	b := bytes.Buffer{}
	b.WriteString("hello")
	fmt.Fprintf(&b, " world")
	b.WriteTo(os.Stdout)
	fmt.Println(b.String())

	buf := bytes.NewBufferString("R29waGVycyBydWxlIQ==")
	dec := base64.NewDecoder(base64.StdEncoding, buf)
	io.Copy(os.Stdout, dec)
}
