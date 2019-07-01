package main

import (
	"fmt"
)

func foo(array *[1e6]int) {
	fmt.Println(array)
}

func main() {
	var array [1e6]int

	foo(&array)
}