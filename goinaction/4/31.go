package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3,4}
	fmt.Println(slice)

	newSlice := slice[2:4]
	fmt.Println(newSlice)

	newSlice = append(newSlice, 55)

	fmt.Println(slice)
	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))
	fmt.Println(cap(slice))
}