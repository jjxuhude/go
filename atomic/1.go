package main

import (
	"errors"
	"fmt"
	"sync/atomic"
)

type ConCurrentArray interface {
	Set(index uint32, elem int) (err error)
	Get(index uint32) (elem int, err error)
	Len() uint32
}

type ConCurrentArrayIns struct {
	len   uint32
	value atomic.Value
}

func NewConCurrentArray(len uint32) ConCurrentArray {
	array := ConCurrentArrayIns{}
	array.len = len
	array.value.Store(make([]int, array.len))
	return &array
}

func (array *ConCurrentArrayIns) checkIndex(index uint32) error {
	if index >= array.len {
		return fmt.Errorf("下标越界")
	}
	return nil
}

func (array *ConCurrentArrayIns) checkValue() error {
	v := array.value.Load()
	if v == nil {
		return errors.New("值为空")
	}
	return nil
}

func (array *ConCurrentArrayIns) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	newArray := make([]int, array.len)
	copy(newArray, array.value.Load().([]int))
	newArray[index] = elem
	array.value.Store(newArray)
	return
}

func (array *ConCurrentArrayIns) Get(index uint32) (elem int, err error) {
	if err = array.checkValue(); err != nil {
		return
	}
	if err = array.checkIndex(index); err != nil {
		return
	}
	elem = array.value.Load().([]int)[index]
	return
}

func (array *ConCurrentArrayIns) Len() uint32 {
	return array.len
}

func main() {
	//demo();
}
func demo() {
	i32 := int32(10)
	atomic.AddInt32(&i32, 4)
	atomic.AddInt32(&i32, -10)

	fmt.Println(i32)
	atomic.CompareAndSwapInt32(&i32, 41, 10)
	fmt.Println(i32)
	delta := int32(100)
	func(delta int32) {
		v := atomic.LoadInt32(&i32)
		for {
			if atomic.CompareAndSwapInt32(&i32, v, (v + delta)) {
				break
			}
		}
	}(delta)
	fmt.Println(i32)
	atomic.StoreInt32(&i32, 200)
	fmt.Println(i32)
	old := atomic.SwapInt32(&i32, 300)
	fmt.Println(old, i32)
}
