package main

import (
	"fmt"
)

func main() {
	slice1 := make([]string, 5)	
	slice2 := make([]int, 3, 5)

	fmt.Println(slice1)
	fmt.Println(slice2)

	//slice3 := make([]string, 5, 3)

	slice4 := []string {"Red", "Blue", "Green"}
	fmt.Println(slice4)

	slice5 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice5)

	slice6 := []string{99: "h"}
	fmt.Println(slice6)

	// 数组和切片
	array1 := [4]string{"h", "w", "x", "x"}
	fmt.Println(array1)

	slice9 := []string{"h", "w", "x", "x"}
	fmt.Println(slice9)


	var slice10 []string
	fmt.Println(slice10)

	if (nil == slice10) {
		fmt.Println("nil slice")
	}

	slice11 := make([]int, 0)
	slice12 := []int{}
	fmt.Println(slice11)
	fmt.Println(slice12)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice12))
}