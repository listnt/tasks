package main

import (
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(40) + 20                                // длина от 20 до 60
	a := new(big.Int).SetBytes([]byte(RandStringBytes(n))) // случайный bigint
	n = rand.Intn(40) + 20
	b := new(big.Int).SetBytes([]byte(RandStringBytes(n))) // случайный bigint

	// Проводим операции
	c := new(big.Int)
	c.Add(a, b)
	d := new(big.Int)
	d.Sub(a, b)
	e := new(big.Int)
	e.Mul(a, b)
	g := new(big.Int)
	g.Div(a, b)

	fmt.Printf("A  :%v\nB  :%v\nAdd:%v\nSub:%v\nMul:%v\nDiv:%v", a, b, c, d, e, g)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//генерация случайных строк
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
