package main 

import (
	"fmt"
	"log"
	"net/http"
)
//http://127.0.0.1:9901/xxxxxxxx
//Url path %q
///xxxxxxxx
func main () {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:9901", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Url path %q\n", r.URL.Path)
}