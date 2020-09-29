package main

import (
	"time"
	"fmt"
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("hello.json")
	fmt.Println(post)
	if err != nil {
		t.Error("Wrong", err)
	}
	if post.Id != 1 {
		t.Error("wrong id", post.Id)
	}
	if post.Content != "Hello world" {
		t.Error("wrong content", post.Content)
	}
}

func TestEncode(t *testing.T) {
	t.Skip("skip ttttt")
}

// func TestFail(t *testing.T) {
// 	t.Error("xxx")
// }

func TestLongRunning(t *testing.T) {
	if testing.Short() {
		t.Skip("xxx")
	}
	time.Sleep(10 * time.Second)
}