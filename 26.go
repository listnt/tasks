package main

import "fmt"

func main() {
	s := "snow dog sun"
	run := make([]rune, len(s))
	for i := len(s) - 1; i > -1; i-- {
		run[len(s)-1-i] = rune(s[i])
	}
	s = string(run)
	fmt.Printf("%s\n", s)
}
