package datafile

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

func TestNewFile(t *testing.T) {
	path := filepath.Join(os.TempDir(), "1205.txt")
	t.Logf("测试文件路径：%s", path)
	datalen := uint32(3)
	data := Data{'A', 'B', 'C'}
	f, _ := NewDataFile(path, datalen)
	defer f.Close()
	t.Run("测试写", func(t *testing.T) {
		var wg sync.WaitGroup
		var n = 6
		wg.Add(n)
		for i := 0; i < 3; i++ {
			go func() {
				defer wg.Done()
				testWrite(f, data, t)
			}()
		}
		for j := 0; j < 3; j++ {
			go func() {
				defer wg.Done()
				testRead(f, t)
			}()
		}
		wg.Wait()
	})
}

func testWrite(f DataFile, data Data, t *testing.T) {
	wsn, _ := f.Write(data)
	fmt.Println("wsn:", wsn)
}

func testRead(f DataFile, t *testing.T) {
	rsn, data, _ := f.Read()
	fmt.Printf("RSN:%d,%c\n", rsn, data)
}
