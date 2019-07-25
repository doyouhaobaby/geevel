package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3,4,4}
	fmt.Println(slice)

	slice[3] = 10000
	fmt.Println(slice)

	slice2 := slice[1:3]
	fmt.Println(slice2)

	slice2[1] = 9999
	fmt.Println(slice)
	fmt.Println(slice2)

	// slice2[88] = 44 
	// panic: runtime error: index out of range

	
}