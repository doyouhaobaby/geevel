package main

import (
	"fmt"
)

func main() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	fmt.Println(array1)
	fmt.Println(array2)

	*array2[0] = "Red"
	*array2[0] = "Blue"
	*array2[0] = "Green"
	fmt.Println(array2)

	array1 = array2
	fmt.Println(array1)
}