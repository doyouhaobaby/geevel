package main 

import (
	"fmt"
)

func main() {
	array := [5]int{10, 20, 30, 40, 50}

	array[3] = 999

	fmt.Println(array)
}
