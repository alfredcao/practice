package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "    Hello world abc    \n"
	result := strings.Replace(str, "world", "you", 5)
	fmt.Println("Replace", result)

	count := strings.Count(str, "l")
	fmt.Println("Count : ", count)

	hasPrefix := strings.HasPrefix(str, "    Hello")
	fmt.Println("HasPrefix : ", hasPrefix)
	hasSuffix := strings.HasSuffix(str, "abc    \n")
	fmt.Println("HasSuffix : ", hasSuffix)

	index := strings.Index(str, "l")
	fmt.Println("Index : ", index)

	lastIndex := strings.LastIndex(str, "l")
	fmt.Println("LastIndex : ", lastIndex)

	result = strings.Repeat(str, 3)
	fmt.Println("Repeat : ", result)

	result = strings.ToLower(str)
	fmt.Println("ToLower : ", result)

	result = strings.ToUpper(str)
	fmt.Println("ToUpper : ", result)

	result = strings.TrimSpace(str)
	fmt.Println("TrimSpace : ", result)

	result = strings.Trim(str, " ")
	fmt.Println("Trim : ", result)

	result = strings.TrimLeft(str, "H \n")
	fmt.Println("TrimLeft : ", result)

	result = strings.TrimRight(str, " \n")
	fmt.Println("TrimRight : ", result)

	slice := strings.Fields(str)
	fmt.Println("Fields : ")
	for _, e := range slice {
		fmt.Println(e)
	}

	slice = strings.Split(str, "l")
	fmt.Println("Split : ")
	for _, e := range slice {
		fmt.Println(e)
	}

	result = strings.Join(slice, "l")
	fmt.Println("Join : ", result)

	result = strconv.Itoa(100)
	fmt.Println("Itoa : ", result, "Type : ", fmt.Sprintf("%T", result))

	number, _ := strconv.Atoi(result)
	fmt.Println("Atoi : ", number, "Type : ", fmt.Sprintf("%T", number))

}
