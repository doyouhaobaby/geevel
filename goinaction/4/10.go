package main

import (
	"fmt"
)

func main() {
	var array [4][2]int
	fmt.Println(array)

	array1 := [4][2]int{{10, 1}, {5, 5}, {5, 8}, {7, 7}}
	fmt.Println(array1)

	array2 := [4][2]int{1: {10, 1}, 3:{7, 7}}
	fmt.Println(array2)
}