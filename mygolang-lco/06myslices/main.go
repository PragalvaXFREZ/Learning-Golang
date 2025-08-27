package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Mango", "Random"}
	fmt.Printf("Type of fruit list is %T \n", fruitList)

	fruitList = append(fruitList, "Banana")
	fmt.Println("\n")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 222
	highScores[1] = 277
	highScores[2] = 235
	highScores[3] = 224

	highScores = append(highScores, 555, 666, 444)

	sort.Ints(highScores)

	fmt.Println(highScores)
}
