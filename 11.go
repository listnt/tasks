package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

func main() {
	arr1 := make([]int, N)
	arr2 := make([]int, N)
	rand.Seed(time.Now().Unix())
	for i := 0; i < N; i++ {
		arr1[i] = rand.Int() % 100
		arr2[i] = rand.Int() % 100
	}
	res := Intersection(arr1, arr2)
	fmt.Println(res, len(res))
}

func Intersection(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	var out []int
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]]++
	}
	for i := 0; i < len(arr2); i++ {
		if m[arr2[i]] > 0 {
			m[arr2[i]]--
			out = append(out, arr2[i])
		}
	}
	return out
}
