package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"time"
)

type BestPlayer struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

func main() {
	bytes, err := PlayerOfTheMatch()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bytes))

	var bestPlayer BestPlayer
	err = json.Unmarshal(bytes, &bestPlayer)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v is Player of the Match! He scored %g goals\n", bestPlayer.Name, bestPlayer.Score)

	DateOfTheMatch()
}

func findMax(x float64, y float64) float64 {
	return math.Max(x, y)
}

func PlayerOfTheMatch() ([]byte, error) {
	m := make(map[string]float64)

	m["CristanoRonaldo"] = 3.0
	m["LeoMessi"] = 2.0

	p1 := m["CristanoRonaldo"]
	p2 := m["LeoMessi"]

	bestPlayer := BestPlayer{}
	bestPlayer.Score = findMax(p1, p2)

	for k, v := range m {
		if v == bestPlayer.Score {
			bestPlayer.Name = k
			break
		}
	}

	jsonBytes, err := json.Marshal(bestPlayer)
	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func DateOfTheMatch() {
	presentTime := time.Now() 
	fmt.Printf("Match date: %v\n", presentTime.Format("01-02-2006 15:00 Thursday"))
}
