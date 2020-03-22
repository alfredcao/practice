package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//testMemoryStore()
	testFileStore()
}

var (
	PostById     = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)
)

func testFileStore() {
	// use package ioutil write&read file
	data := []byte("Hello World!")
	err := ioutil.WriteFile("data1", data, 0777)
	if err != nil {
		panic(err)
	}
	data1, _ := ioutil.ReadFile("data1")
	fmt.Println("read from data1, content :", string(data1))

	// use File write&read file
	file, err := os.Create("data2")
	if err != nil {
		panic(err)
	}
	num, _ := file.Write(data)
	fmt.Printf("write to data2 %d bytes\n", num)

	file2, _ := os.Open("data2")
	data2 := make([]byte, len(data))
	file2.Read(data2)
	fmt.Println("read from data2, content :", string(data2))

}

/**
内存存储
*/
func testMemoryStore() {
	post1 := Post{1, "Hello!", "user1"}
	post2 := Post{2, "你好!", "user2"}
	post3 := Post{3, "What's your name?", "user3"}
	post4 := Post{4, "Fine!", "user1"}
	storeInMemory(&post1)
	storeInMemory(&post2)
	storeInMemory(&post3)
	storeInMemory(&post4)

	fmt.Println(PostById[1])
	for _, v := range PostByAuthor["user1"] {
		fmt.Println(v)
	}
}

func storeInMemory(post *Post) {
	PostById[post.Id] = post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], post)
}

type Post struct {
	Id      int
	Content string
	Author  string
}
