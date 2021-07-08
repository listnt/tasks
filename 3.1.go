package main

import (
	"os"
	"strconv"
)

//constants: N-array size, numberOfGourutines=number threads
var N = 10000
var numberOfGourutines = 40

//driver code
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
	arr := make([]int, N)          // массив и его наполнение
	for i := 0; i < N; i++ {
		arr[i] = (i + 1) * 2
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

// функция вычисляет сумму (элементы слайса в степени 2)
func MyFunc(input []int, done chan int) {
	var sum int
	for i := 0; i < len(input); i++ {
		sum += input[i] * input[i]
	}
	done <- sum
}
