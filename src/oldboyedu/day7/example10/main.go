package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//testStruct()
	//testMap()
	testSlice()
}

func testSlice() {
	jsonStr := `[{"userName":"user1","nickName":"xiaoming","age":18,"birthday":"1989-02-23","email":"mahuateng@qq.com","sex":"男","phone":"110"}]`
	var users []User
	err := json.Unmarshal([]byte(jsonStr), &users)
	if err != nil {
		fmt.Println("反序列化错误！", err)
	} else {
		fmt.Println(users)
	}
}

func testMap() {
	jsonStr := `{"userName":"user1","nickName":"xiaoming","age":18,"birthday":"1989-02-23","email":"mahuateng@qq.com","sex":"男","phone":"110"}`
	var user1 map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &user1)
	if err != nil {
		fmt.Println("反序列化错误！", err)
	} else {
		fmt.Println(user1)
	}
}

func testStruct() {
	jsonStr := `{"userName":"user1","nickName":"xiaoming","age":18,"birthday":"1989-02-23","email":"mahuateng@qq.com","sex":"男","phone":"110"}`
	var user1 User
	err := json.Unmarshal([]byte(jsonStr), &user1)
	if err != nil {
		fmt.Println("反序列化错误！", err)
	} else {
		fmt.Println(user1)
	}
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
