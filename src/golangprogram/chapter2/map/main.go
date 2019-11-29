package main

import "fmt"

func main() {
	var personDB map[string]Person
	//personDB = make(map[string] Person)
	//personDB["1234"] = Person{"12345", "Tom", "Room 100"}
	//personDB["1"] = Person{"1", "Jack", "Room 101"}

	personDB = map[string]Person{
		"1234": {"12345", "Tom", "Room 100"},
		"1":    {"1", "Jack", "Room 101"},
	}
	fmt.Println("personDB大小：", len(personDB))

	person, ok := personDB["1234"]
	if ok {
		fmt.Println("找到了ID为1234的人：", person)
	} else {
		fmt.Println("没找到了ID为1234的人")
	}

	delete(personDB, "1234")
	fmt.Println("删除ID为1234的人！")

	person1, ok1 := personDB["1234"]
	if ok1 {
		fmt.Println("找到了ID为1234的人：", person1)
	} else {
		fmt.Println("没找到了ID为1234的人")
	}
}

type Person struct {
	ID      string
	Name    string
	Address string
}
