package main

import "fmt"

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	fmt.Println(v)
	justString = v[:100]
}

func main() {
	someFunc()
}

func createHugeString(a int) string {
	s := make([]byte, a)
	return string(s)
}
