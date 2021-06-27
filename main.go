package main

import "fmt"

func main() {
	var d1 [32]int32
	var d2 [32]int32
	for i := 0; i < 32; i++ {
		d1[i] = 1
		d2[i] = 1
	}

	var sum2 int32 = 0
	sum2 = VDotProdAVX2(d1[:], d2[:])
	fmt.Println("Sum: ", sum2)
}
