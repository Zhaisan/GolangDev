package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

//func split(sum int) (int, int) {
//	temp := sum * 4 / 9
//  return temp, sum - temp
//}

func main() {
	fmt.Println(split(17))
}
