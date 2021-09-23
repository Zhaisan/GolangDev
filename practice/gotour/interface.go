package main

import "fmt"

type animal interface {
	makeSound()
}

type cat struct {}
type dog struct {}

func (c *cat) makeSound(){
	fmt.Println("meow!")
}

func (d *dog) makeSound(){
	fmt.Println("guf guf!")
}
func main() {
	var c, d animal = &cat{}, &dog{}
	c.makeSound()
	d.makeSound()
}