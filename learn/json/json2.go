package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	jsonFile, err := os.Open("hello.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}

	defer jsonFile.Close()
	decoder := json.NewDecoder(jsonFile)
	var post Post
	for {
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error", err)
			return
		}
	}
	fmt.Println(post)
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
