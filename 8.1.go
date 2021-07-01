package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//просим ввести данные, обрабатываем ошибки
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

	// вычисляем по модулю 64 - количество битов в х64 архитектуре
	bit = bit % 64
	var mask int
	var res = num   //оригинальное значение заносим в переменную результата
	if bit_v != 0 { //если необходимо установить бит в 1,
		//то производим оперцию num |(ИЛИ)) 00001000(маска условна)
		mask = bit_v << bit
		res |= mask
	} else { // если необходимо устовить бит в 0,
		//то производим оперцию с маской  num &(И) 111111011111(маска условна)
		mask = ^(1 << bit)
		res = res & mask
	}
	//выводим результаты
	fmt.Printf("original:0x%x , value:%d\nmasked:0x%x , value:%d",
		num, num, res, res)
}
