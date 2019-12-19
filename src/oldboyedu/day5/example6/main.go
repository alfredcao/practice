package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Student struct {
	Id   string
	Name string
	Age  int
}

type StudentSlice []Student

func (p StudentSlice) Len() int {
	return len(p)
}

func (p StudentSlice) Less(i, j int) bool {
	return p[i].Name < p[j].Name
}

func (p StudentSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var stuSli StudentSlice
	for i := 1; i < 10; i++ {
		stu := Student{
			Id:   fmt.Sprintf("360729%d", rand.Int()),
			Name: fmt.Sprintf("stu%d", rand.Intn(100)),
			Age:  rand.Intn(20),
		}

		stuSli = append(stuSli, stu)
	}
	fmt.Println("排序前：")
	for _, v := range stuSli {
		fmt.Println(v)
	}

	fmt.Println("排序后：")
	sort.Sort(stuSli)
	for _, v := range stuSli {
		fmt.Println(v)
	}
}
