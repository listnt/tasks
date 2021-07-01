package main

import (
	"fmt"
	"math"
)

// точка и вектор имеют примерно одинаковые свойства
// поэтому длина и расстояние между точками имеет смысл
type Point interface {
	Dest(a Point) float64
	Len() float64
	get() point
}

//с маленькой буквы не импортируется
type point struct {
	x float64 //значение по оси х
	y float64 //значение по оси у
}

//расстояние между точками
func (p *point) Dest(a Point) float64 {
	ap := a.get()
	return (&point{ap.x - p.x, ap.y - p.y}).Len()
}

//Длина от точки до начала координат
func (p *point) Len() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// Геттер, прмечание - возвращает копию, а не адрес на структуру,
// т.е. меняй не меняй полученную структуру ничего не изменишь.
// Да и методов у копии нет, и не импортируется в другие пакеты
func (p *point) get() point {
	return *p
}

func NewPoint(x float64, y float64) Point {
	// Инкапсулируем внутри пакета путем создания совместимого интефеса.
	// А снаружи пакета он кроме методов Dest и Len ничего и не увидет,
	// get не увидит ибо с маленькой буквы

	return &point{x, y}
}

func main() {
	a := NewPoint(12, 13)
	b := NewPoint(12, -7)
	fmt.Printf("dist between points: %f", a.Dest(b)) // Assert.equal(20,a.Dest(b))
}
