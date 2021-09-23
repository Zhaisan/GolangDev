package main

import "fmt"

type animal interface{
	walker
	runner
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type cat struct {}

func (c *cat) walk() {
	fmt.Println("cat is walking")
}

func (c *cat) run() {
	fmt.Println("cat is running")
}

func main() {
	var c animal = &cat{}
	c.walk()
	c.run()



}