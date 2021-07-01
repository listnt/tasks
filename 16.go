package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1 //локальная переменная
		n++    //изменение локальной переменной
	}
	fmt.Println(n) //"глобальная" переменная
}
