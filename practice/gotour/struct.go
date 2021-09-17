package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)
	fmt.Println(Vertex{1, 2})

}
