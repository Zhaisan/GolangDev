package main

import (
	"fmt"
	"time"
)

func main(){
	go fmt.Println("Hello from goroutine!")

	fmt.Println("Hello from main()")

	time.Sleep(time.Millisecond)
}