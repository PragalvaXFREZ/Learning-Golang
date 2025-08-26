package main

import (
	"fmt"
)

func main() {
	var username string = "Pragalva"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("Variable is of type: %T \n", smallValue)

	var smallFloat float32 = 255.323156
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values and some aliases

	var anotherVar int
	fmt.Println((anotherVar))
	fmt.Printf("Variable is of type: %T \n", anotherVar)

	//implicit type

	var sebiste = "pragalva.com"
	fmt.Println(sebiste)

	//no var style

	numberofUser := 30000
	fmt.Println((numberofUser))

}
