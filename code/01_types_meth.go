package main

import (
	"fmt"
	"math"
)

type Complex struct {
	real    float64
	imag    float64
}

func (cn *Complex) modulus() float64 {
	sum := math.Pow(cn.real, 2) + math.Pow(cn.imag, 2)
	return math.Sqrt(sum)
}

func main() {
	n := new(Complex)
	n.real = 3
	n.imag = 4
	fmt.Printf("%f\n", n.modulus())
}
