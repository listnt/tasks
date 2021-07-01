package main

import (
	"fmt"
	"math/rand"
)

const N = 10000

//код проверки функции быстрой сортировки
func main() {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = rand.Int() % 10000
	}
	arr = quicksort(arr, 0, N-1)
	fmt.Println(arr[:100])
}

// очередная быстрая сортировка
func quicksort(arr []int, l int, h int) []int {
	if len(arr) < 2 {
		return arr
	}
	size := h - l + 1
	stack := make([]int, size)
	top := -1
	top = top + 1
	stack[top] = l
	top = top + 1
	stack[top] = h
	for top >= 0 {
		h = stack[top]
		top = top - 1
		l = stack[top]
		top = top - 1
		p := partition(arr, l, h)
		if p-1 > l {
			top = top + 1
			stack[top] = l
			top = top + 1
			stack[top] = p - 1
		}
		if p+1 < h {
			top = top + 1
			stack[top] = p + 1
			top = top + 1
			stack[top] = h
		}
	}
	return arr
}

func partition(arr []int, l int, h int) int {
	i := (l - 1)
	x := arr[h]
	for j := l; j < h; j++ {
		if arr[j] <= x {
			i = i + 1
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[h] = arr[h], arr[i+1]
	return i + 1
}
