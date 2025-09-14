package main

import (
	"fmt"
)

func main() {
	//no inheritance in golang; No super or parent concepts

	pragalva := User{"Pragalva", "sapkotapragalva@gmail.com", true, 20}
	fmt.Println(pragalva)
	fmt.Printf("Pragalva details are: %+v\n", pragalva)

	fmt.Printf("Name is %v and email is  %v. ", pragalva.Name, pragalva.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
