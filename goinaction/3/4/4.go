package main

import (
	"fmt"
	myfmt "mylib/fmt" // 仅仅演示 mylib 不存在
)

func main() {
	fmt.Println("stxxx")
	myfmt.Println("xxxxxx")
}