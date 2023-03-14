package main

import "fmt"

func main() {
	fmt.Println("Enter a number")
	var n1 int
	fmt.Scanln(&n1)

	fmt.Println("Enter a number")
	var n2 int
	fmt.Scanln(&n2)

	fmt.Println("Enter what action is to be taken with the entered values")
	fmt.Println("Enter 1 to 'add'")
	fmt.Println("Enter 2 to 'subtract'")
	fmt.Println("Enter 3 to 'multiply'")
	fmt.Println("Enter 4 to 'divide'")
	fmt.Println("Enter 5 to 'print name'")
	var n3 int
	fmt.Scanln(&n3)

	switch n3 {
	case 1:
		fmt.Println(n1 + n2)
	//	break
	case 2:
		fmt.Println(n1 - n2)
	//	break
	case 3:
		fmt.Println(n1 * n2)
	//	break
	case 4:
		fmt.Println(n1 / n2)
	//	break
	case 5:
		fmt.Println("Enter your first name : ")
		var fname string
		fmt.Scanln(&fname)

		fmt.Println("Enter your last name : ")
		var lname string
		fmt.Scanln(&lname)

		fmt.Println(fname + " " + lname)
	}
}
