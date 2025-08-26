package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to my pizza place")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the rating for our pizza")

	//comma ok syntax

	input, _ := reader.ReadString('\n') //till when do you want to read? till enter so '\n'
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of rating is : %T", input)
}
