package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkUnique(str string) bool {
	var arr [256]bool
	for i := 0; i < len(str); i++ {
		if arr[int(str[i])] == true {
			return false
		}
		arr[int(str[i])] = true
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if checkUnique(text) {
		os.Stdout.Write([]byte("Unique"))
	} else {
		os.Stdout.Write([]byte("NOT unique"))
	}
}
