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

//намного проще чем задача 4, потому что я решил не париться с обработкой
//сигнала прерывания
func main() {
	chanin := make(chan int, 2)  //канал входных данных
	chanout := make(chan int, 2) //канал выходных данных
	var wg sync.WaitGroup        //вэйтгруппа
	quit := make(chan bool, 2)   //канал для оповещении об окончании работы

	//увеличиваем счетчик на 2 и создаем 2 горутины
	wg.Add(2)
	go dataReader(chanin, quit, &wg)
	go dataProcessing(chanin, chanout, quit, &wg)

	//канал в который приходит сигнал прерывания
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	// цикл обработки каналов
	for {
		select {
		case t := <-chanout: //вывод результатов
			fmt.Println(t)
		case <-osSignals: //сигнал прерывания
			quit <- true
			quit <- true
			fmt.Println("[!!!]Exiting")
			wg.Wait()
			return
		}
	}
}

//генерация случайных данных в канал
func dataReader(outchan chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case <-quit:
			return
		case <-ticker.C: // по таймеру генерируем данные
			outchan <- rand.Int() % 100
		}
	}
}

//обработка данных из входного канала и последующая запись в выходной канал
func dataProcessing(inchan chan int, outchan chan int, quit chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-quit:
			return
		case t := <-inchan: // Если есть данные в канале
			if t%2 == 0 { // Проверка на четность
				outchan <- HighlySophisticatedMath(t) // мало ли, вдруг математика поменяется
			}
		}
	}
}

//очень сложная для понимания математика
func HighlySophisticatedMath(a int) int {
	return a
}
