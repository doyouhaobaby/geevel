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

	var array1 [2][2]int
	array1 = array
	fmt.Println(array1)

	var array3 [2]int = array[1]
	fmt.Println(array3)

	var value int = array[1][0]
	fmt.Println(value)
}