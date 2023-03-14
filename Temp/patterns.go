package main

import "fmt"

func main() {
	var i int
	var j int
	var count = 1
	var patterns int

	fmt.Println("Enter the pattern you want to print : ")
	fmt.Println("Enter 1 to print triangle with 90 degree angle at left")
	fmt.Println("Enter 2 to print triangle with 90 degree angle at right")
	fmt.Println("Enter 3 to print triangle with 90 degree angle at centre")
	fmt.Println("Enter 4 to print triangle upside down with 90 degree on left")
	fmt.Println("Enter 5 to print triangle upside down with 90 degree on right")
	fmt.Println("Enter 6 to print triangle upside down with 90 degree on centre")
	fmt.Scanln(&patterns)

	switch patterns {
	case 1:
		for i = 0; i < 5; i++ {
			fmt.Println("* ")
			for j = 0; j < count; j++ {
				fmt.Print("*")
				count++
			}
		}
	case 2:
		for i = 0; i < 5; i++ {
			for j = 0; j < 5; j++ {
				fmt.Println("*")
			}
		}
	}
}
