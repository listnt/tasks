package main

import (
	"math/rand"
	"sync"
	"time"
)

var lock = sync.RWMutex{}
var quit chan bool

func main() {
	quit = make(chan bool, 2)
	arr := make([]int, 100)
	go write(arr)
	go read(arr)
	go read(arr)
	go read(arr)
	go read(arr)
	time.Sleep(2 * time.Second)
	quit <- true
	quit <- true
	quit <- true
	quit <- true
	quit <- true

	time.Sleep(1 * time.Second)
}
func write(arr []int) {
	for {
		select {
		case <-quit:
			return
		default:
			lock.Lock()
			arr[rand.Int()%100] = rand.Int()
			lock.Unlock()
		}
	}
}
func read(arr []int) {
	for {
		select {
		case <-quit:
			return
		default:
			lock.RLock()
			_ = arr[rand.Int()%100]
			lock.RUnlock()
		}
	}
}
