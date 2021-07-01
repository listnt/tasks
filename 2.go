package main

import (
	"fmt"
	"os"
	"strings"
)

//constants: N-array size, numberOfGourutines=number threads
var N = 563
var numberOfGourutines = 4

//driver code
func main() {
	//объявление пременных
	chunkSize := int(N / numberOfGourutines) // размер чанка данных
	var chansize = 4
	if i := N - chunkSize*numberOfGourutines; i > 0 { //если есть отсаток от деления
		chansize++
	}
	ch := make(chan int, chansize) // канал для оповещении об окончании работы
	arr := make([]int, N)          // массив и его наполнение
	for i := 0; i < N; i++ {
		arr[i] = (i + 1) * 2
	}
	for i := 0; i < numberOfGourutines; i++ {
		//нарезаем массив на чанки одинакового размера
		go MyFunc(arr[chunkSize*i:chunkSize*(i+1)], ch)
	}
	if i := N - chunkSize*numberOfGourutines; i > 0 {
		//Если есть остаток от деления, то выполняется дополнительная горутина на этот остаток
		go MyFunc(arr[chunkSize*numberOfGourutines:], ch)
	}
	for i := 0; i < numberOfGourutines; i++ {
		//ждем завершения
		<-ch
	}
	//записываем в поток
	os.Stdout.Write([]byte(strings.Trim(fmt.Sprint(arr[0:20]), "[]")))
}

//функция по обновлению данных в массиве
func MyFunc(input []int, done chan int) {
	for i := 0; i < len(input); i++ {
		input[i] *= input[i]
	}
	done <- 1
}
