package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	post, err := decode("hello.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(post)
}

func decode(file string) (post Post, err error) {
	jsonFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	//for {
	err = decoder.Decode(&post)
	// if err == io.EOF {
	// 	break
	// }
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	//}
	//fmt.Println(post)
	return
}

func unmarshal(filename string) (post Post, err error) {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	json.Unmarshal(jsonData, &post)
	return
}

// post.
type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

// auth.
type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// comment.
type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}
