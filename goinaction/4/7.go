package main

import (
	"fmt"
)

func main() {	
	var array1 [5]string
	
	array2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}
	
	array1 = array2

	fmt.Println(array1)
}