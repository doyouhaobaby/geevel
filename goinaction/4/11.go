package main

import (
	"fmt"
)

func main() {
	var array [2][2]int
	fmt.Println(array)

	array[0][0] = 1
	array[0][1] = 2
	fmt.Println(array)
	array[1][0] = 5
	array[1][1] = 3
	fmt.Println(array)
}