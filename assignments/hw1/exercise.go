package main

import (
	"fmt"
	"math"
)

const inf = 1e-13

func Sqrt(x float64) float64 {
	z, y := 1.0, 0.0
	for math.Abs(z - y) > inf {
		y = z
		z -= (z * z - x) / (2 * z)
		fmt.Println(z)
	}
	return z


}

func main() {
	a := Sqrt(3)
	b := math.Sqrt(3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("Difference:", a - b)
}
