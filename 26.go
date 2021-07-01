package main

import "fmt"

func main() {
	s := "snow dog sun"
	//переворот строки
	run := make([]rune, len(s))
	for i := len(s) - 1; i > -1; i-- {
		run[len(s)-1-i] = rune(s[i]) //заносим значения с конца строки
	}
	s = string(run)
	fmt.Printf("%s\n", s)
}
