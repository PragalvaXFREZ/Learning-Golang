package main

import "fmt"

func main() {
	// var ptr *int

	// fmt.Println("Value of pointer is: ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println(": ", ptr)  //ADDRESS
	fmt.Println(": ", *ptr) //absolute value

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)
}
