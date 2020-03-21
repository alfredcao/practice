package main

import "fmt"

func main() {
	testMemoryStore()
}

var (
	PostById     = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)
)

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
