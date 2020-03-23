package main

import (
	"bytes"
	"encoding/csv"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//testMemoryStore()
	//testFileStore()
	//testCsvFileStore()
	testGobFileStore()
}

var (
	PostById     = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)
)

func testGobFileStore() {
	post := Post{1, "Hello!", "user1"}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(post)
	if err != nil {
		panic(err)
	}
	fileName := "gobdata"
	err = ioutil.WriteFile(fileName, buffer.Bytes(), 0777)
	if err != nil {
		panic(err)
	}

	fmt.Println("write to file success")

	raw, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	buffer1 := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer1)
	post1 := new(Post)
	err = decoder.Decode(&post1)
	if err != nil {
		panic(err)
	}

	fmt.Println("read from file :", post1)
}

/**
CSV文件存储
*/
func testCsvFileStore() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("process failed, err :", err)
		}
	}()
	csvFileName := "posts.csv"

	// 创建并写入CSV文件
	csvFile, err := os.Create(csvFileName)
	if err != nil {
		panic(err)
	}

	defer csvFile.Close()
	csvWrite := csv.NewWriter(csvFile)
	postList := []Post{
		{1, "Hello!", "user1"},
		{2, "你好!", "user2"},
		{3, "What's your name?", "user3"},
		{4, "Fine, Thank you!", "user1"},
	}

	for _, post := range postList {
		err = csvWrite.Write([]string{strconv.Itoa(post.Id), post.Content, post.Author})
		if err != nil {
			panic(err)
		}
	}

	csvWrite.Flush()
	fmt.Println("写入CSV文件成功！")

	// 从CSV文件读取内容
	file, err := os.Open(csvFileName)
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println("CSV文件内容如下：")
	for _, record := range records {
		columnCount := len(record)
		for index, columnValue := range record {
			if strings.Contains(columnValue, ",") {
				columnValue = fmt.Sprintf(`"%s"`, columnValue)
			}
			if index == columnCount-1 {
				fmt.Printf(columnValue + "\n")
			} else {
				fmt.Printf(columnValue + ",")
			}
		}
	}
}

/**
文件存储
*/
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
	defer file.Close()
	num, _ := file.Write(data)
	fmt.Printf("write to data2 %d bytes\n", num)

	file2, _ := os.Open("data2")
	defer file2.Close()
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
