package main

import "fmt"

func main() {
	var a, b int
	a = 10
	b = 5
	fmt.Printf("before swap a:%d, b:%d\n", a, b)
	a, b = b, a //свап
	fmt.Printf("after swap a:%d, b:%d\n", a, b)

}
