package main

import (
	"fmt"
	"time"
)

func main() {
	ch := UserTicker()
	//time.Sleep(20 * time.Second)
	ch <- true
	close(ch)
}

func UserTicker() chan bool {
	ticker := time.NewTicker(5 * time.Second)

	stopChan := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				fmt.Println("Ticker2....")
			case stop := <-stopChan:
				if stop {
					fmt.Println("Ticker2 Stop")
					return
				}
			}
		}
	}(ticker)

	return stopChan
}
