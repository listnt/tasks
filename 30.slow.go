package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	i := 2

	// Убираем элемент i из массива
	copy(a[i:], a[i+1:]) // Смещаем массив на один влево
	a[len(a)-1] = ""     // затираем последниий элемент
	a = a[:len(a)-1]     // Получаем новый слайс
	// !!Примечание cap на один больше len
	fmt.Println(a) // [A B D E]
}
