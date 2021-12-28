package main

import (
	"fmt"
	"go.uber.org/dig"
	"math"
)

// 2x^2 + 4x - 32 = 0
func main() {
	var a float64
	var b float64
	var c float64
	a = 2
	b = 4
	c = -32
	fmt.Println(FindRoots(a, b, c))
}

func FindRoots(a float64, b float64, c float64) (float64, float64) {
	D := FindDiscriminant(a, b, c)
	x1 := -b + math.Sqrt(D)/(2*a)
	x2 := -b - math.Sqrt(D)/(2*a)
	return x1, x2
}

func FindDiscriminant(a float64, b float64, c float64) float64 {
	D := b*b - 4*a*c
	return D
}

func BuildContainer() *dig.Container {
	container := dig.New()
	_ = container.Provide(NewConfig)
	_ = container.Provide(NewDB)
	_ = container.Provide(NewServer)
	return container
}
