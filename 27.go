package main

import (
	"fmt"
	"strings"
)

func reverse(numbers []string) []string {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i] //свапаем левую и правую строки
	}
	return numbers
}

func main() {
	var str = "snow dog sun"
	// обожаю one-liners
	fmt.Println(strings.Join(reverse(strings.Split(str, " ")), " "))
}
