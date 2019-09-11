package main

import "testing"
func TestCalculate(t *testing.T) {
	if Calculate(2) !=2.2 {
	t.Error("Expected it to equal 2.2")
	}
}
