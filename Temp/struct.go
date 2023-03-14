package main

import "fmt"

type Temp struct {
	no    int
	fname string
	lname string
}

func main() {
	var person Temp

	fmt.Println("Enter your no. : ")
	fmt.Scanln(&person.no)
	fmt.Println("Enter your first name : ")
	fmt.Scanln(&person.fname)
	fmt.Println("Enter your last name : ")
	fmt.Scanln(&person.lname)

	fmt.Println("Your number is : ", person.no)
	fmt.Println("Your number name is : ", person.fname+" "+person.lname)
	//fmt.Println("Your number last name is : ", person.lname)
}
