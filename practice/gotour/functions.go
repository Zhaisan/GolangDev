package main

import "fmt"

func add(x int, y int) int {      // or add(x, y int) int{ ... }
	return x + y
}

func main() {
	xWithY := add(1, 2)
	fmt.Println(xWithY)
	fmt.Println(add(23, 12))

}


