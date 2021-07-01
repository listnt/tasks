package main

import (
	"math/rand"
	"sync"
	"time"
)

//объявляем переменные
var m = map[int]int{1: 1, 2: 2} //мапа для записи
var lock = sync.RWMutex{}       //мютекс
var cm sync.Map                 //мапа с синхронизацией
var quit chan bool              //канал для оповещении о закрытии

func main() {
	quit = make(chan bool, 2) //создаем канал
	// три писателя, один читатель
	// плохие книги не читают
	// эта особенна плоха
	go write()
	go write()
	go write()
	go read()

	//ждем пока отработает некоторое время
	time.Sleep(2 * time.Second)

	quit <- true //начинаем закрывать
	quit <- true
	quit <- true
	quit <- true
	time.Sleep(1 * time.Second) //подождем пока закроется
}

//писатель
func write() {
	for {
		select {
		case <-quit: //кейс закрыть
			return
		default:
			lock.Lock() //блокируем для чтения записи
			m[rand.Int()%100] = rand.Int() % 100
			lock.Unlock()
		}
	}
}
func read() {
	for {
		select {
		case <-quit:
			return
		default:
			lock.RLock() //блокируем для чтения, т.е никто не сможет
			//записывать пока читатели(их может быть несколько, см задание 21) не закончат чтение
			_ = m[rand.Int()%100]
			lock.RUnlock()
		}
	}
}
