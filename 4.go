package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
)

//воркер
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1
	for j := range jobs {
		//опять работа
		//some actions
		fmt.Println("worker", id, "started  job:", j)
		time.Sleep(time.Second)
		results <- j * 2
		fmt.Println("worker", id, "finished job:", j)
	}
}

var numWorkers int

//driver code
func main() {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(os.Stdin) //ридер потока
	//чтение количества воркеров обработка ошибок
	fmt.Print("Enter number of workers: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	var err error
	numWorkers, err = strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results := make(chan int, numWorkers) //канал для результатов
	jobs := make(chan int, numWorkers)    //канал для работы
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ { //создаем воркеры
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	syncchan := make(chan bool, 1)

	go MainCycle(jobs, results, syncchan)
	wg.Wait()        // Когда Все воркеры отработали
	syncchan <- true // Посылаем сигнал завершения
}

func MainCycle(jobs chan int, results chan int, syncchan chan bool) {
	//канал для сигнала прерывания
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	//главный цикл
	for {
		select {
		case jobs <- rand.Int() % 100: //генерация данных
		case t := <-results: //чтение результатов
			fmt.Println(t)
		case <-osSignals: //обработка сигнала прерывания
			close(jobs)
			fmt.Println("[!!!]Got signal, Processing")
			//дожидаемся оставшиеся результаты
			for {
				select {
				case t := <-results: //читаем результат
					fmt.Println(t)
				case <-syncchan: //когда все воркеры отработали оставшиеся работы - завершаем
					fmt.Println("[!!!]Exiting")
					return
				}
			}
		}
	}
}
