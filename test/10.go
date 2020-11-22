package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://www.hao123.com/", strings.NewReader("name=jack.xu1"))
	rep, _ := client.Do(req)
	data, _ := ioutil.ReadAll(rep.Body)
	fmt.Println(string(data))
	defer rep.Body.Close()

}
