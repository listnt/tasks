package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 100

func main() {
	//объявляем переменные
	arr1 := make([]int, N)
	arr2 := make([]int, N)
	rand.Seed(time.Now().Unix())
	//наполняем случайными значениями
	for i := 0; i < N; i++ {
		arr1[i] = rand.Int() % 100
		arr2[i] = rand.Int() % 100
	}
	//выводим результат пересечения
	res := Intersection(arr1, arr2)
	fmt.Println(res, len(res))
}

// вычисляем пересечение двух слайсов через мапу, сохраняет дубликаты
// Пример: (2,3,6,2,5) * (2,2,6,12,32) -> (2,2,6)
func Intersection(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	var out []int
	//для каждого значения в первом слайсе, увеличиваем соответствующее значение в мапе
	for i := 0; i < len(arr1); i++ {
		m[arr1[i]]++
	}

	//для каждого значения в втором слайсе. Если в мапе есть это значение,
	//то значит оно есть и в первом слайсе -> элемент добаляется в выходной слайс
	for i := 0; i < len(arr2); i++ {
		if m[arr2[i]] > 0 {
			m[arr2[i]]--
			out = append(out, arr2[i])
		}
	}
	return out
}
