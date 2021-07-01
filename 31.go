package main

import (
	"fmt"
	"math"
)

type Point interface {
	Dest(a Point) float64
	Len() float64
	get() *point
}

type point struct {
	x float64
	y float64
}

func (p *point) Dest(a Point) float64 {
	return (&point{a.get().x - p.x, a.get().y - p.y}).Len()
}

func (p *point) Len() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *point) get() *point {
	return p
}

func NewPoint(x float64, y float64) Point {
	return &point{x, y}
}

func main() {
	a := NewPoint(12, 13)
	b := NewPoint(-5, -7)
	fmt.Printf("dist between points: %f", a.Dest(b))
}
