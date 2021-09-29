package main

import "fmt"

func main() {
	message := make(chan string , 2)
	message <- "hello"

	fmt.Println(<-message)
}