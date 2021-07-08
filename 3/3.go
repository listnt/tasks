package main

import (
	"os"
	"strconv"
)

var N = 10000
var numberOfGourutines = 40

func main() {
	//объявление пременных
	chunkSize := int(N / numberOfGourutines) // размер чанка данных
	var chansize = numberOfGourutines
	if i := N - chunkSize*numberOfGourutines; i > 0 {
		//если есть отсаток от деления, то размер канала на 1 больше
		//для дополнительной горутины
		chansize++
	}
	ch := make(chan int, chansize) // канал для оповещении об окончании работы
	arr := make([]int64, N)        // массив и его наполнение
	for i := 0; i < N; i++ {
		arr[i] = int64(i+1) * 2
	}
	for i := 0; i < numberOfGourutines; i++ {
		//нарезаем массив на слайсы одинакового размера
		go MyFunc(arr[chunkSize*i:chunkSize*(i+1)], ch)
	}
	if i := N - chunkSize*numberOfGourutines; i > 0 {
		//Если есть остаток от деления, то выполняется дополнительная горутина на этот остаток
		go MyFunc(arr[chunkSize*numberOfGourutines:], ch)
	}
	var sum int
	for i := 0; i < chansize; i++ {
		sum += <-ch
	}
	//запись в поток
	os.Stdout.Write([]byte("final sum: " + strconv.Itoa(sum) + "\n"))
}

func MyFunc(input []int64, done chan int) {
	sum := int(VecSum(input))
	//fmt.Println(sum)
	done <- sum
}
