package main

import (
	"fmt"
)

const eurotodollars = 1.1
const eurotopounds = 0.89

func main() {
	for{

		var choice int
		var euro float64

		fmt.Println("Welcome")
		fmt.Println("Enter 1 for euro to dollars")
		fmt.Println("Enter 2 for euro to pounds")
		fmt.Println("Enter 3 for the function to terminate!!")

		fmt.Println("Enter choice:")
		fmt.Scanf("%d", &choice)

		if choice == 3 {
			fmt.Println("Goodbye")
			break;
		}
	
		fmt.Println("Enter the amount of euro:")
		fmt.Scanf("%f", &euro)

		
		if choice == 1 {
			dollars := euro * eurotodollars
			fmt.Println("dollars", dollars)
		} else if choice == 2 {
			pound := euro * eurotopounds
			fmt.Println("pounds", pound)
		} 
}
}
