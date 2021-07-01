package main

import (
	"fmt"

	"internal/task31"
)

func main() {
	a := task31.NewPoint(12, 13)
	b := task31.NewPoint(12, -7)
	fmt.Printf("dist between points: %f", a.Dest(b)) // Assert.equal(20,a.Dest(b))
}
