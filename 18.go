package main

import "fmt"

func someAction(v []int8, b int8) { //получаем (адрес массива, длину, вместимость), новое значение
	//и заносим их в локальные копии
	v[0] = 100
	v = append(v, b) //локальная переменная
}

func main() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6) //передаем (адрес массива, длину, вместимость), новое значение
	fmt.Println(a)
}