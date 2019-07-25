package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3,4,4}
	fmt.Println(slice)

	newSlice := slice[1:3]
	fmt.Println(newSlice)

	newSlice = append(newSlice, 55)

	fmt.Println(slice)
	fmt.Println(newSlice)
}