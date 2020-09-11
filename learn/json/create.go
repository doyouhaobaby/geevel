package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

func main() {
	post := Post{
		Id:      1,
		Content: "Hello",
		Author: Author{
			Id:   2,
			Name: "saxx",
		},
		Comments: []Comment{
			Comment{
				Id:      44,
				Content: "hell",
				Author:  "44",
			},
			Comment{
				Id:      3,
				Content: "hell",
				Author:  "44",
			}},
	}

	output, err := json.MarshalIndent(&post, "", "    ")
	if err != nil {
		fmt.Println("Error marshaling to JSON:", err)
		return
	}

	err = ioutil.WriteFile("create.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}
