package main

import "testing"

func TestCalculate(t *testing.T) {
	if findMax(1.0, 2.0) != 2 {
		t.Error("Expected Max(1.0 , 2.0) is equal to 2")
	}
}