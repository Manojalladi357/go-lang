package main

import (
	"fmt"
	"math"
)

func main() {
	var no int

	fmt.Println("enter a number : ")
	fmt.Scanln(&no)

	if no < 2 {
		fmt.Println("Enter a number greater than 2")
		return
	}

	sqr_root := int(math.Sqrt(float64(no)))

	for i := 2; i <= sqr_root; i++ {
		if no%i == 0 {
			fmt.Println("Not a Prime number.")
			return
		}
	}

	fmt.Println("Prime number.")
	return
}
