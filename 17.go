package main

import (
	"fmt"
)

func main() {
	var ty int
	var k interface{}
	fmt.Println("Enter Type of variable\n1. int\n2. string\n3. bool\n4. chan")
	fmt.Scanf("%d\n", &ty)
	switch ty {
	case 1:
		k = 2
	case 2:
		k = "asd"
	case 3:
		k = true
	case 4:
		k = make(chan int, 1)
	default:
		fmt.Println("Wrong choise, try again")
	}

	switch k.(type) {
	case int:
		fmt.Print("Your choise is int")
	case string:
		fmt.Print("Your choise is string")
	case bool:
		fmt.Print("Your choise is bool")
	case chan int:
		fmt.Print("Your choise is chan")
	}
}
