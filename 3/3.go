package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var N = 1000
var numberOfGourutines = 4

func main() {

	st := time.Now()
	ch := make(chan int, 4)
	arr := make([]int64, N)
	for i := 0; i < N; i++ {
		arr[i] = int64((i + 1) * 2)
	}
	chunkSize := int(N / numberOfGourutines)
	for i := 0; i < numberOfGourutines; i++ {
		go MyFunc(arr[chunkSize*i:chunkSize*(i+1)], ch)
	}
	if i := N - chunkSize*numberOfGourutines; i > 0 {
		go MyFunc(arr[chunkSize*numberOfGourutines:], ch)
	}
	var sum int
	for i := 0; i < numberOfGourutines; i++ {
		sum += <-ch
	}
	os.Stdout.Write([]byte(strconv.Itoa(sum)))

	fmt.Printf("\n%d", (time.Now().Sub(st)).Milliseconds())
}

func MyFunc(input []int64, done chan int) {
	a := len(input)
	sum := int(VecSum(input[:int(a/4)*4]))
	for i := int(a/4) * 4; i < a; i++ {
		sum += int(input[i] * input[i])
	}
	done <- sum
}
