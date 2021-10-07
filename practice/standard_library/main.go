package main

import (
	"flag"
	"fmt"
	"lectures/std/math"
)

var (
	a, b int
)

func init() {
	flag.IntVar(&a, "a", 0, "first argument for adder")
	flag.IntVar(&b, "b", 0, "second argument for adder")
}
func main() {
	flag.Parse()
	added := math.Adder(a, b)
	fmt.Printf("The sum of %d and %d is %d\n", a, b, added)
}