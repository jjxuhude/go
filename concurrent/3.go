package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	sigRec1 := make(chan os.Signal)
	sigs1 := []os.Signal{syscall.SIGINT, syscall.SIGQUIT}
	signal.Notify(sigRec1, sigs1...)

	sigRec2 := make(chan os.Signal)
	sigs2 := []os.Signal{syscall.SIGINT}
	signal.Notify(sigRec2, sigs2...)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for sig := range sigRec1 {
			fmt.Printf("sigRec1:%s\n", sig)
		}
		wg.Done()
	}()

	go func() {
		for sig := range sigRec2 {
			fmt.Printf("sigRec2:%s\n", sig)
		}
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	signal.Stop(sigRec1)
	close(sigRec1)

	wg.Wait()

}
