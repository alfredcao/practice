package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testStruct()
	//testInt()
	//testMap()
	testSlice()
}

type User struct {
	UserName string `json:"userName"`
	NickName string `json:"nickName"`
	Age      int    `json:"age"`
	Birthday string `json:"birthday"`
	Email    string `json:"email"`
	Sex      string `json:"sex"`
	Phone    string `json:"phone"`
}

func testInt() {
	var age = 18

	jsonByteArr, err := json.Marshal(age)
	if err != nil {
		fmt.Println("json序列化失败！", err)
	} else {
		fmt.Println(string(jsonByteArr))
	}
}

func testMap() {
	var userMap = make(map[string]*User)
	user1 := &User{
		UserName: "user1",
		NickName: "xiaoming",
		Age:      18,
		Birthday: "1989-02-23",
		Email:    "mahuateng@qq.com",
		Sex:      "男",
		Phone:    "110",
	}
	userMap["user1"] = user1
	user2 := &User{
		UserName: "user2",
		NickName: "xiaohong",
		Age:      20,
		Birthday: "1999-04-12",
		Email:    "mayun@qq.com",
		Sex:      "男",
		Phone:    "119",
	}
	userMap["user2"] = user2

	jsonByteArr, err := json.Marshal(userMap)
	if err != nil {
		fmt.Println("json序列化失败！", err)
	} else {
		fmt.Println(string(jsonByteArr))
	}
}

func testSlice() {
	var userSlice = make([]*User, 2)
	user1 := &User{
		UserName: "user1",
		NickName: "xiaoming",
		Age:      18,
		Birthday: "1989-02-23",
		Email:    "mahuateng@qq.com",
		Sex:      "男",
		Phone:    "110",
	}
	userSlice[0] = user1
	user2 := &User{
		UserName: "user2",
		NickName: "xiaohong",
		Age:      20,
		Birthday: "1999-04-12",
		Email:    "mayun@qq.com",
		Sex:      "男",
		Phone:    "119",
	}
	userSlice[1] = user2

	jsonByteArr, err := json.Marshal(userSlice)
	if err != nil {
		fmt.Println("json序列化失败！", err)
	} else {
		fmt.Println(string(jsonByteArr))
	}
}

func testStruct() {
	user1 := &User{
		UserName: "user1",
		NickName: "xiaoming",
		Age:      18,
		Birthday: "1989-02-23",
		Email:    "mahuateng@qq.com",
		Sex:      "男",
		Phone:    "110",
	}

	jsonByteArr, err := json.Marshal(user1)
	if err != nil {
		fmt.Println("json序列化失败！", err)
	} else {
		fmt.Println(string(jsonByteArr))
	}
}
