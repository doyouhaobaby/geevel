package main 

import (
	"fmt"
	"os"
)

//go run echo.go es ff xxxx
func main () {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = "."
	}
	fmt.Println(s)
}