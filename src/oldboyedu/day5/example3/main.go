package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	stu1 := Student{
		Name:  "stu1",
		Age:   18,
		score: 100,
	}

	data, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("json格式化异常！", err)
	} else {
		fmt.Println(string(data))
	}
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	score int    `json:"score"`
	Tag   string `json:"tag,omitempty"`
}
