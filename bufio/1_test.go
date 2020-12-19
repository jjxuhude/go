package main

import (
	"fmt"
	"testing"
)

func TestLine(t *testing.T) {
	data, _ := line()
	fmt.Println("订单数量:", len(data))

}
func TestWrite(t *testing.T) {
	writer()

}

func BenchmarkLine(b *testing.B) {
	b.ResetTimer()
	line()
}
