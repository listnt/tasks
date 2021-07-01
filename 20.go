package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice[0] = "b"
		slice = append(slice, "a")

		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
