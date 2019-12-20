package main

import "fmt"

func main() {
	//typeAssert1("hehe")
	var map1 = make(map[string]string, 10)
	map1["key"] = "value"
	typeAssert2(1, 8.99, true, "hello", Student{Name: "stu1", Age: 18}, &Student{Name: "stu2", Age: 20}, map1)
}

func typeAssert1(x interface{}) {
	target, ok := x.(int)
	if !ok {
		fmt.Println("convert failed!")
		return
	}
	target += 3
	fmt.Println(target)

}

func typeAssert2(items ...interface{}) {
	for i, v := range items {
		switch v.(type) {
		case int, int8, int16, int32, int64:
			fmt.Printf("%d param type is int, value is %v\n", i, v)
		case float32, float64:
			fmt.Printf("%d param type is float, value is %v\n", i, v)
		case string:
			fmt.Printf("%d param type is string, value is %v\n", i, v)
		case bool:
			fmt.Printf("%d param type is bool, value is %v\n", i, v)
		case nil:
			fmt.Printf("%d param type is nil, value is %v\n", i, v)
		case map[string]string:
			fmt.Printf("%d param type is map[string]string, value is %v\n", i, v)
		case Student:
			fmt.Printf("%d param type is Student, value is %v\n", i, v)
		case *Student:
			fmt.Printf("%d param type is *Student, value is %v\n", i, v)

		}
	}

}

type Student struct {
	Name string
	Age  int
}
