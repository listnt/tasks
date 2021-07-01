package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	chanin := make(chan int, 1)
	chanout := make(chan int, 1)
	var wg sync.WaitGroup
	quit := make(chan bool, 2)

	wg.Add(2)
	go dataReader(chanin, quit, &wg)
	go dataProcessing(chanin, chanout, quit, &wg)

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case t := <-chanout:
			fmt.Println(t)
		case <-osSignals:
			quit <- true
			quit <- true
			fmt.Println("[!!!]Exiting")
			wg.Wait()
			return
		}
	}
}

func dataReader(outchan chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-quit:
			return
		case <-ticker.C:
			outchan <- rand.Int() % 100
		}
	}
}

func dataProcessing(inchan chan int, outchan chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			return
		case t := <-inchan:
			if t%2 == 0 {
				outchan <- HighlySophisticatedMath(t)
			}
		}
	}
}

func HighlySophisticatedMath(a int) int {
	return a
}
