package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) 
	fmt.Println("My favourite number is", rand.Intn(100))
}