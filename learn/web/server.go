package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8082",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	//log.Println("Method is:", r)
	log.Println("Method is:", r.Method)
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	case "PUT":
		err = handlePut(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	log.Println(r.URL.Path)
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	log.Println(path.Base(r.URL.Path))
	log.Println(id)
	if err != nil {
		return
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Conetent-Type", "application/json")
	w.Write(output)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()
	log.Println(r.Form)
	for k, v := range r.Form {
		log.Println("key:", k)
		log.Println("value:", v)
	}
	log.Println(r.Form["content"])
	// len := r.ContentLength
	// body := make([]byte, len)
	// r.Body.Read(body)
	log.Println("form value test", r.FormValue("content"))
	var post Post
	// post.Content = "xxx"
	// post.Author = "y2342"
	// //	log.Println(body)
	// json.Unmarshal(body, &post)
	log.Println(post)
	post.Content = r.FormValue("content")
	post.Author = r.FormValue("author")
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("path is:", path.Base(r.URL.Path))
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	err = post.delete(id)
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	fmt.Println("path is:", path.Base(r.URL.Path))
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	fmt.Println(id)
	var post Post
	post.Id = id
	post.Content = r.FormValue("content")
	post.Author = r.FormValue("author")
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
