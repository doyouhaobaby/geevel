package main

import (
	"os"
	"encoding/json"
	"fmt"
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

	jsonFile, err := os.Create("create2.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}
