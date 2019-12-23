package main

import (
	"fmt"
	"reflect"
)

func main() {
	stu := Student{
		Name:  "stu",
		Age:   18,
		Score: 90,
	}
	fmt.Println(stu)
	structReflect(&stu)
	fmt.Println(stu)
}

func structReflect(s interface{}) {
	val := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	kind := val.Kind()
	if !(kind == reflect.Ptr && val.Elem().Kind() == reflect.Struct) {
		panic("请传入结构体！")
	}

	numField := val.Elem().NumField()
	fmt.Printf("共有 %d 个字段\n", numField)
	fmt.Println("每个字段的json标签如下：")
	for i := 0; i < numField; i++ {
		field := t.Elem().Field(i)
		tag := field.Tag
		fmt.Printf("%d : %s\n", i, tag.Get("json"))
	}
	val.Elem().FieldByName("Name").SetString("stu1")
	val.Elem().FieldByName("Age").SetInt(20)
	val.Elem().FieldByName("Score").SetInt(98)

	numMethod := val.Elem().NumMethod()
	fmt.Printf("共有 %d 个方法\n", numMethod)
	for i := 0; i < numMethod; i++ {
		val.Elem().Method(i).Call(nil)
	}
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (stu Student) Print() {
	fmt.Printf("Name : %s, Age : %d, Sex : %d\n", stu.Name, stu.Age, stu.Score)
}
