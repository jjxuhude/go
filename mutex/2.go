package main

import (
	"errors"
	"os"
	"sync"
)

type Data []byte

type DataFile interface {
	Rsn() int64
	Wsn() int64
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	datalen uint32
	wmutex  sync.Mutex
	rmutex  sync.Mutex
	woffset int64
	roffset int64
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.datalen)
	df.rmutex.Unlock()

	rsn = offset / int64(df.datalen)

	return
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	return
}

func (df *myDataFile) DataLen() uint32 {
	return df.datalen
}

func (df *myDataFile) Rsn() int64 {
	return df.roffset
}

func (df *myDataFile) Wsn() int64 {
	return df.woffset
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.Close()
}

func newDataFile(path string, datalen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if datalen == 0 {
		return nil, errors.New("获取的数据长度不能为0.")
	}
	df := &myDataFile{f: f, datalen: datalen}
	return df, nil
}
func main() {

}
