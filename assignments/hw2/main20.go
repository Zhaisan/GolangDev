package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64
const inf = 1e-13

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number : %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
    	return 0, ErrNegativeSqrt(x)
    }

    z, y := 1.0, 0.0
	for math.Abs(z - y) > inf {
		y = z
		z -= (z * z - x) / (2 * z)
		fmt.Println(z)
	}
	return z, nil
}

func main() {
    fmt.Println(Sqrt(3))
    fmt.Println(Sqrt(-3))
}