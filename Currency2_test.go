package main

import "testing"
func TestCalculate(t *testing.T) {
	if Calculate(2) !=1.78 {
	t.Error("Expected it to equal 1.78")
	}
}
