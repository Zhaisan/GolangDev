package main

import (
	"fmt"
	"sync"
)


func main() {
	c1 := check(1, 2, 3, 4, 5)
	c2 := check(6, 7, 8, 9, 10)
	c3 := check(11, 12, 13, 14, 15)
	c4:= check(16, 17, 18, 19, 20)
	c5 := check(21, 22, 23, 24, 25)
	for v := range merge(c1, c2, c3, c4, c5) {
		fmt.Println(v)
	}
}


func merge(cs ... <-chan int) <-chan int {

	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, i := range cs {
		go func(i <-chan int) {
			for v := range i {
				ch <- v
			}
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func check(vs ... int) <-chan int {

	j := make(chan int)
	go func() {
		for _, v := range vs {
			j <- v
		}
		close(j)
	}()
	return j
}