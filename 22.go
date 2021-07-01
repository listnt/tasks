package main

import (
	"fmt"
	"math/rand"
)

const N = 10000

func main() {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Int() % 10000
	}
	arr = quicksort(arr)
	fmt.Println(arr[:100])
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := (left + right) / 2

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
