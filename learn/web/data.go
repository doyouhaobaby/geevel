package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/go-learn-bbs?charset=utf8")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connect to MySQL success")
	}
}

func retrieve(id int) (post Post, err error) {
	post = Post{}
	fmt.Println("retrieve id ", id)
	err = Db.QueryRow("select id,content,author from posts where id = ?", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post) create() (err error) {
	statement := "insert into posts(content, author) values(?, ?)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	ret, err := stmt.Exec(post.Content, post.Author)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if lastInsertID, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", lastInsertID)
		post.Id = int(lastInsertID)
	}
	fmt.Println(post)
	return
}

func (post *Post) delete(id int) (err error) {
	_, err = Db.Exec("delete from posts where id = ?", id)
	return
}

func (post *Post) update() (err error) {
	_, err = Db.Exec("update posts set content = ?, author = ? where id = ?", post.Content, post.Author, post.Id)
	return
}
