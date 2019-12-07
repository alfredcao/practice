package main

import (
	"encoding/json"
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	log.Println("log base files : ", "./var/log")
	fmt.Println(filepath.Join("/var/log", "app.log"))

	// struct tag
	user := &User{"chronos", "pass"}
	//s := reflect.TypeOf(user).Elem()
	//for i := 0; i < s.NumField(); i++ {
	//	fmt.Println(s.Field(i).Tag.Get("user"))
	//}
	content, err := json.Marshal(user)
	if err == nil {
		fmt.Println(string(content))
	}
}

type User struct {
	Name   string `json:"name"`
	Passwd string
}
