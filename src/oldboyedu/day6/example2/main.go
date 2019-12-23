package main

import (
	"fmt"
	"reflect"
)

func main() {
	//a := 1
	//stu := Student{
	//	Name: "stu",
	//	Age:  19,
	//}
	//reflectTest(stu)
}

type Student struct {
	Name string
	Age  int
}

func reflectTest(a interface{}) {
	t := reflect.TypeOf(a)
	fmt.Printf("Type is %v\n", t)
	fmt.Printf("Type Name is %s\n", t.Name())
	fmt.Printf("Type String is %s\n", t.String())
	fmt.Printf("Type Kind is %v\n", t.Kind())
	fmt.Printf("Type Size is %v\n", t.Size())

	v := reflect.ValueOf(a)
	fmt.Printf("Value is %v\n", v)
	fmt.Printf("Value String is %v\n", v.String())
	fmt.Printf("Value Kind is %v\n", v.Kind())
	fmt.Printf("Value IsValid is %v\n", v.IsValid())
	fmt.Printf("Value Field(0) is %v\n", v.Field(0))
	fmt.Printf("Value Type is %v\n", v.Type())
}
