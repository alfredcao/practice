package main

import "fmt"

func main() {
	testMap()
}

func testMap() {
	var map1 map[string]int = make(map[string]int, 10)
	var map2 map[string]map[string]string = make(map[string]map[string]string, 10)
	map2["zhangsan"] = make(map[string]string, 10)
	map2["zhangsan"]["id"] = "001"
	map2["zhangsan"]["password"] = "1qaz@WSX"
	map2["lisi"] = make(map[string]string, 10)
	map2["lisi"]["id"] = "002"
	map2["lisi"]["password"] = "1qaz@WSX"

	var map3 = map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}

	fmt.Println("遍历map1：")
	for k, v := range map1 {
		fmt.Println(k, v)
	}

	fmt.Println("遍历map2：")
	for k, v := range map2 {
		fmt.Println(k)
		for k1, v1 := range v {
			fmt.Println("\t", k1, v1)
		}
	}

	fmt.Println("遍历map3：")
	for k, v := range map3 {
		fmt.Println(k, v)
	}

	_, ok := map2["zhangsan"]
	if ok {
		fmt.Println("map2中存在zhangsan!")
	} else {
		fmt.Println("map2中不存在zhangsan!")
	}
	_, ok1 := map2["wangwu"]
	if ok1 {
		fmt.Println("map2中存在wangwu!")
	} else {
		fmt.Println("map2中不存在wangwu!")
	}

	delete(map3, "key1")

	fmt.Println("删除key1后遍历map3：")
	for k, v := range map3 {
		fmt.Println(k, v)
	}
	fmt.Println("删除key1后map3长度：", len(map3))
}
