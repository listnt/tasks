package main

import (
	"bufio"
	"fmt"
	"os"
)

// проверка на уникальность символов в строке
func checkUnique(str string) bool {
	var arr [256]bool // В ASCII всего 256 символов
	for i := 0; i < len(str); i++ {
		if arr[int(str[i])] == true { // Если какой то символ встретился второй раз
			return false
		}
		arr[int(str[i])] = true
	}
	return true
}

func main() {
	//чтение о обработка ошибок
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//проверка на уникальность символов в тексте
	if checkUnique(text) {
		os.Stdout.Write([]byte("Unique"))
	} else {
		os.Stdout.Write([]byte("NOT unique"))
	}
}
