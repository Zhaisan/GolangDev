package main

import (
	"encoding/json"
	"fmt"
	"math"
	"time"
)


func findMax(x float64, y float64) float64 {
	return math.Max(x, y)
}

func PlayerOfTheMatch() {

	sx := make(map[string]float64)

	sx["CristanoRonaldo"] = 1.0
	sx["LeoMessi"] = 2.0

	res1 := sx["CristanoRonaldo"]
	res2 := sx["LeoMessi"]
	score := findMax(res1, res2)

	jsonStr, err := json.Marshal(sx)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	} else {
		fmt.Println(string(jsonStr))
	}

	if sx["CristanoRonaldo"] == score {
		fmt.Printf("Cristano is Player of the Match! He scored %g goals\n", sx["CristanoRonaldo"])
	} else{
		fmt.Printf("Leo is Player of the Match! He scored %g goals\n", sx["LeoMessi"])
	}

}

func DateOfTheMatch() {
	
	presentTime := time.Now()
	fmt.Printf("Match date: %v\n", presentTime.Format("01-02-2006 15:00 Thursday"))
}




func main() {
	PlayerOfTheMatch()
	DateOfTheMatch()
}