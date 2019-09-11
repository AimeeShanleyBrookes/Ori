package main

import (
	"fmt"
)

func Calculate(euro float64) (result float64) {
	result = euro * 1.1
	return result
}

func main(){
	fmt.Println("Go Testing Tutorial")
	result := Calculate(2)
	fmt.Println(result)
}
