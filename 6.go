package main

import (
	"sync"
	"time"
)

func main() {

	//1-ый путь: путь закрытия канала
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	go func() {
		for {
			foo, ok := <-ch
			if !ok {
				println("done")
				wg.Done()
				return
			}
			println(foo)
		}
	}()
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	wg.Wait()
	//плох тем, что есть операция в бесконечном цикле.
	//Вероятность съесть все ресурсы

	//2-ой путь: путь посылки сигнала
	quit := make(chan bool)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				//somestuff
			}
		}
	}()
	time.Sleep(1 * time.Second)
	quit <- true
	//3-ий путь: не путю
}
