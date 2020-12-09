package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	len := 3
	arr := NewConCurrentArray(uint32(len))
	t.Logf("arr:%v", arr)
	arr.Set(0, 111)
	arr.Set(1, 222)
	arr.Set(2, 333)
	t.Logf("arr:%v", arr)
	elem, _ := arr.Get(2)
	t.Logf("elem 2:%v", elem)
}
