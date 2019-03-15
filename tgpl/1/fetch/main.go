package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
//go run main.go https://baidu.com
func main () {
	for _,url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch err %v \n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch reading %s %v", url, err)
		}
		fmt.Printf("%s", b)
	}
}