package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	go parseURL("http://cisco.com")
	parseURL("http://youtube.com")

	fmt.Printf("Parsing completed. Time Elapsed: %.2f seconds\n", time.Since(t).Seconds())
}

func parseURL(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("Parsing <%s> - Step - %d - Latency %d ms\n", url, i + 1, latency)
	}
}