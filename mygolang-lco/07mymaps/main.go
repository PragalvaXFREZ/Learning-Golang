package main

import (
	"fmt"
)

func main() {

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("list of all lanugages : ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}

}

// %v = value
