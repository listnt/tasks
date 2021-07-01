package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BitMaskWrap(num int, bit int, bit_v bool) int {
	return int(BitMask(int64(num), int64(bit), bit_v))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	num, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Enter number of Bit: ")
	text, _ = reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	bit, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Print("Enter value: ")
	text, _ = reader.ReadString('\n')
	text = strings.Trim(text, "\r\n")
	bit_v, err := strconv.Atoi(text)
	if err != nil || (bit_v != 0 && bit_v != 1) {
		if err == nil {
			err = errors.New("wrong value. Bit can be zero or one")
		}
		fmt.Println(err)
		os.Exit(1)
	}
	bit_value := (bit_v != 0)
	fmt.Printf("original:0x%x , value:%d\nmasked:0x%x , value:%d",
		num, num, BitMaskWrap(num, bit, bit_value),
		BitMaskWrap(num, bit, bit_value))
}
