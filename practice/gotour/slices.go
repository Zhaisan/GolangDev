package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s, len(s), cap(s))

	s = append(s, 11)
	s = append(s, 14)
	fmt.Println(s, len(s), cap(s))

	s = append(s, 123)
	fmt.Println(s, len(s), cap(s))


}
