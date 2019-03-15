package main 

import (
	"fmt"
	"os"
	"strings"
)

//go run echo.go es ff xxxx
func main () {
	fmt.Println(strings.Join(os.Args[1:], "-"))
}