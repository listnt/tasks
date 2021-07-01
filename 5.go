package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//писатель в канал
func writer(writechan chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker1 := time.NewTicker(100 * time.Millisecond) //тикер, когда тикнет запишет в канал
	for {
		select {
		case <-done: //если получил сигнал на завершение, то завершает
			fmt.Println("Writer stopped")
			ticker1.Stop()
			return
		case <-ticker1.C: //если тикнет, то запишет в канал
			writechan <- (rand.Int() % 1000)
		}
	}
}

//читатель канала
func reader(readchan chan int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-done: //если получил сигнал на завершение работы
			fmt.Println("Reader stopped")
			return
		case t := <-readchan: //если есть данные в канале, то прочитает
			fmt.Printf("Readed: %d\n", t)
		}
	}
}

//драйвер код
func main() {
	//объявляем перееменные
	done := make(chan bool, 2)
	chanel := make(chan int, 1)
	var wg sync.WaitGroup

	timer1 := time.NewTimer(2 * time.Second) //запускаем таймер на N секунд
	wg.Add(2)                                //увеличиваем счетчик на 2
	go writer(chanel, done, &wg)             //запускаем горутины
	go reader(chanel, done, &wg)
	<-timer1.C //ждем таймер

	done <- true //посылаем сигнал о завершении
	done <- true

	wg.Wait() //ждем пока закончит
}
