package main

import "fmt"


//НЕ БОЛЬШЕ ДВУХ ВХОДНЫХ ПАРАМЕТРОВ!!!
func swap(word1, word2 string) (string, string) {
	return word2, word1
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}