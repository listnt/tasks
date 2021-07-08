package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "АААBBCCCDDEEE"
	ar := []rune(a)
	b := ""
	var count int = 1
	for i := 0; i < len(ar)-1; i++ {
		if ar[i] == ar[i+1] {
			count++
		} else {
			if count == 1 {
				b = b + string(ar[i])
			} else {
				b = b + strconv.Itoa(count) + string(ar[i])
			}
			count = 1
		}
		if i == len(ar)-2 {
			b = b + strconv.Itoa(count) + string(ar[i])
		}
	}
	fmt.Printf("%s", b)
}
