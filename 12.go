package main

import "fmt"

func update(p *int) { // сохраняем адрес переменной в локальную переменную p
	b := 2
	p = &b // заносим в локальную переменную адрес b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p) // выводим результат разыменовывания указателя p - a
	update(p)       //передаем адрес переменной a
	fmt.Println(*p)
}
