package main

import "fmt"


type greeter interface{
	greet(string) string
}

type russian struct {}
type kazakh struct {}

func (r *russian) greet(name string) string {
	return fmt.Sprintf("Привет, %s", name)
}

func (k *kazakh) greet(name string) string {
	return fmt.Sprintf("Салем, %s", name)
}

func sayHello(g greeter, name string) {
	fmt.Println(g.greet(name))
}
func main() {
	sayHello(&russian{}, "Костян")
	sayHello(&kazakh{}, "Жасик")
}