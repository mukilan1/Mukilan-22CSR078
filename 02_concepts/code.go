package main

import "fmt"

func checkNumber(x int) {
	if x > 0 {
		fmt.Println("Positive")
	} else if x < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}
}

func main() {
	checkNumber(5)  // Output: Positive
	checkNumber(-3) // Output: Negative
	checkNumber(0)  // Output: Zero
}
