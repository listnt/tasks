package main

import (
	"math/rand"
	"sync"
	"time"
)

var lock = sync.RWMutex{} //мютекс
var quit chan bool

func main() {
	quit = make(chan bool, 2)
	arr := make([]int, 100)

	//один писатель, три читателя
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

//писатель
func write(arr []int) {
	for {
		select {
		case <-quit:
			return
		default:
			lock.Lock() //блокиреут для чтения и записи
			arr[rand.Int()%100] = rand.Int()
			lock.Unlock()
		}
	}
}

//читатель
func read(arr []int) {
	for {
		select {
		case <-quit:
			return
		default:
			lock.RLock() //блокирует остальные горутины на запись, но не запись
			_ = arr[rand.Int()%100]
			lock.RUnlock()
		}
	}
}
