package main

import "testing"
func TestCalculate(t *testing.T) {
	if Calculate(2) !=3.3 {
	t.Error("Expected it to equal 3.3 ")
	}
}
