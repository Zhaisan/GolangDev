package main

import (
	"fmt"
	//	"math/cmplx"
)

// var (
// 	ToBe   bool       = false
// 	MaxInt uint64     = 1<<64 - 1
// 	z      complex128 = cmplx.Sqrt(-5 + 12i)
// )

func main() {
	helloRussia := "Привет РОССИЯ!"
	//fmt.Println(helloRussia[0])
	//fmt.Println(string(helloRussia[0]))
	for _, myRune := range helloRussia {
		fmt.Println(string(myRune))
	}
}