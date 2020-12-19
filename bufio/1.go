package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var redisDB *redis.Client
var ctx = context.Background()

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

func line() (data []int64, err error) {
	path := filepath.Join(os.TempDir(), "order.csv")
	f, err := os.Open(path)
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	var order string
	getRedis()
	for scanner.Scan() {
		//fmt.Printf("%v\n", scanner.Text())
		order = scanner.Text()
		result, _ := redisDB.RPush(ctx, "aOrder", order).Result()
		//data= append(data,order)
		data = append(data, result)
	}
	return

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

func getRedis() *redis.Client {
	redisDB = redis.NewClient(&redis.Options{
		Addr:     "192.168.145.100:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return redisDB
}

func writer() {
	now := time.Now().Format("2006-01-02")
	now += ".csv"
	path := filepath.Join(os.TempDir(), now)
	f, _ := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0777)
	buf := bytes.NewBuffer([]byte{})
	w := bufio.NewWriter(buf)
	for i := 0; i <= 1000000; i++ {
		i := strconv.Itoa(i + 1)
		w.WriteString(i + ",299869,1912160000282,19121600002826,\"paid\",\"PAID\",0,0,0,43,10000265636,\"Connext Li\",0,0,0,0,\"address\",0,0,0,0,0,0,0,0,0,\"王永乾\",\"上海市\",\"上海市\",\"浦东新区\",\"世纪大道8号国金二期6楼\",15021165981,86,200135,\"WXP\",\"CREATED\",7895231264907287,0,0,0,0,0,0,0,0,1,2931,3450,519,0,0,0,0,0,519,100,0,0,0,0,2931,0,0,0,0,0,0,0,\"2019-12-16 15:10:31\",\"2019-12-16 15:10:46\",\"[]\"\n")
	}
	w.Flush() // 不要忘记刷新
	string := buf.String()
	f.WriteString(string)
	//fmt.Printf("\n%s", buf.String())
}

func toRedis() {
	getRedis()
	redisDB.RPush(ctx, "aOrder", "111").Result()
}

//func main() {
////
////	//reader()
////	//scaner()
////	writer()
//	line()
//	//getRedis()
//	//name,_:=redisDB.RPush(ctx,"aOrder","111").Result()
//	//fmt.Println(name)
////
//}
