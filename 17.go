package main

import (
	"fmt"
)

func main() {
	var ty int
	var var1 interface{}
	fmt.Println("Enter Type of variable\n1. int\n2. string\n3. bool\n4. chan")
	fmt.Scanf("%d\n", &ty)
	//заносим значения различных типов в переменную типа channel{}
	switch ty {
	case 1:
		var1 = 2
	case 2:
		var1 = "asd"
	case 3:
		var1 = true
	case 4:
		var1 = make(chan int, 1)
	default:
		fmt.Println("Wrong choise, try again")
	}

	//свитч кейс для типа var
	switch var1.(type) {
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
