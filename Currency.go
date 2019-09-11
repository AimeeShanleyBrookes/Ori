package main

import (
	"fmt"
)

func Calculate(x float64) (result float64) {
	result = x * 1.1
	return result
}

func main(){
	fmt.Println("Go Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
