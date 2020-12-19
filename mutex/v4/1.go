package datafile

import (
	"errors"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

type Data []byte

type DataFile interface {
	Read() (rsn int64, d Data, err error)
	Write(d Data) (wsn int64, err error)
	RSN() int64
	WSN() int64
	DataLen() uint32
	Close() error
}

type myDataFile struct {
	f       *os.File
	fmutex  sync.RWMutex
	rmutex  sync.Mutex
	wmutex  sync.Mutex
	datalen uint32
	roffset int64
	woffset int64
	rcond   *sync.Cond
}

func NewDataFile(path string, datalen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if datalen == 0 {
		return nil, errors.New("数据长度不能为0")
	}
	df := &myDataFile{f: f, datalen: datalen}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func OpenDataFile(path string, datalen uint32) (DataFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	df := &myDataFile{f: f}
	df.rcond = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}

func (df *myDataFile) Signal() {

}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.roffset)
		if atomic.CompareAndSwapInt64(&df.roffset, offset, (offset + int64(df.datalen))) {
			break
		}
	}

	rsn = offset / int64(df.datalen)
	bytes := make([]byte, df.datalen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()
	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.rcond.Wait()
				continue
			}
			return
		}
		d = bytes
		return
	}

}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.datalen)
	df.wmutex.Unlock()

	wsn = offset / int64(df.datalen)
	var bytes []byte
	if len(d) > int(df.datalen) {
		bytes = d[:df.datalen]
	} else {
		bytes = d
	}

	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	_, err = df.f.Write(bytes)
	df.rcond.Signal()
	return

}

func (df *myDataFile) RSN() int64 {
	//df.rmutex.Lock()
	//defer df.rmutex.Unlock()
	//return df.roffset / int64(df.datalen)
	offset := atomic.LoadInt64(&df.roffset)
	return offset / int64(df.datalen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.datalen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.datalen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}
