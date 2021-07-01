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
func worker(id int, jobs <-chan int, results chan<- int, done chan bool, wg *sync.WaitGroup) {
	defer wg.Done() // уменьшаем счетчик на 1
	for {
		select { //слушаем канал с данными
		case j, ok := <-jobs:
			//этот if нужен, потому, что порой селект читает
			//пустые данные из закрытого канал. Ключевое слово ПУСТЫЕ
			if !ok { //если канал закрыт
				//брекаем switch и ждем пока случай выберет другой кейс
				break
			}
			//some actions
			fmt.Println("worker", id, "started  job:", j)
			time.Sleep(2 * time.Second)
			results <- j * 2
			fmt.Println("worker", id, "finished job:", j)
		case <-done: //если получили сигнал на завершение
			for j := range jobs {
				//обрабатываем оставшиеся работу
				//some actions
				fmt.Println("worker", id, "started  job:", j)
				time.Sleep(time.Second)
				results <- j * 2
				fmt.Println("worker", id, "finished job:", j)
			}
			return
		}
	}
}

//driver code
func main() {
	rand.Seed(time.Now().Unix())
	reader := bufio.NewReader(os.Stdin) //ридер потока
	//чтение количества воркеров обработка ошибок
	fmt.Print("Enter number of workers: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	numWorkers, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results := make(chan int, numWorkers) //канал для результатов
	jobs := make(chan int, numWorkers)    //канал для работы
	done := make(chan bool, numWorkers)   //канал для сообщении об окончании работы
	var wg sync.WaitGroup
	for w := 0; w < numWorkers; w++ { //создаем воркеры
		wg.Add(1)
		go worker(w, jobs, results, done, &wg)
	}
	//канал для сигнала прерывания
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	// конструкция ниже необходима потому, что, когда подается сигнал на завершение работы
	// данные перестают поступать в канал jobs но, там какое-то
	// количество все еще остается, и их надо обработать

	//Но поскольку чтение из канала results происходит в гланом цикле,
	//То может случится, что при получении сигнала окончании и обработки оставшихся работ
	// канал results может оказаться полным, поэтому необходимо
	// 1. дождаться всех результатов 2. Воркеры уменьшают счетчик wg 3. Посылаем сигнал в канал
	syncchan := make(chan bool, 1)
	go func() {
		wg.Wait()
		syncchan <- true
	}()

	//главный цикл
	for {
		select {
		case jobs <- rand.Int() % 100: //генерация данных
		case t := <-results: //чтение результатов
			fmt.Println(t)
		case <-osSignals: //обработка сигнала прерывания

			close(jobs)
			fmt.Println("[!!!]Got signal, Processing")
			for w := 0; w < numWorkers; w++ { //посылаем сигнал о завершении
				done <- true
			}
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
