package main

import (
	"fmt"
	"math/rand"
)

const N = 20

func main() {
	m := make(map[int][]float64) // мапа вида{5:{6,5,9}, -5:{-6,-1,-9}}
	arr := make([]float64, N)
	var t int

	for i := 0; i < N; i++ {
		//наполняем случайными значениями преводим в диапозон (-50,50)
		arr[i] = rand.Float64()
		if arr[i] > 0.5 {
			arr[i] = 0.5 - arr[i]
		}
		arr[i] = 100 * arr[i]
		//вычисляет смещение. -5 для отрицательных, 5 для положительных
		t = -5
		if arr[i] > 0 {
			t = 5
		}

		//заносим значения в мапу
		key := int(arr[i]/10)*10 + t
		m[key] = append(m[key], arr[i])
	}
	fmt.Println(m)
}
