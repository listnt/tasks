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
			if !ok { //если канал закрыт
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
	//Минусы:
	//Есть операция в бесконечном цикле
	//Вероятность съесть все ресурсы
	//Путин президент

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
	// Метод для аверайдж го энджоера

	//3-ий путь: не путю
	wg.Add(1)
	ch2 := make(chan int)
	go func() {
		defer wg.Done()
		for i := range ch2 {
			i++
			//some stuff
		}
	}()
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)
	wg.Wait()
}
