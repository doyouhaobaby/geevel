package main 

import (
	"fmt"
)

func main() {
	array := [5]*int{0: new(int), 1: new(int)}

	*array[1] = 55

	fmt.Println(array)
}
