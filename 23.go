package main

import (
	"fmt"
)

const N = 10000
const val = 200

//код для проверки поиска
func main() {
	arr := make([]int, N)
	for i := 0; i < N; i++ {
		arr[i] = i * 2
	}
	fmt.Printf("found %d at %d", val, BinarySearch(arr, val))
}

//очередной бинарный поиск
func BinarySearch(arr []int, key int) int {
	l := 0
	r := len(arr) - 1
	m := l + (r-l)/2 //на случай переполнения
	for l <= r {
		m = l + (r-l)/2
		if arr[m] == key {
			return m
		}
		if arr[m] < key {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}
