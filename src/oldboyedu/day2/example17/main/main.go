package main

import "fmt"

func main() {
	var str1 = "hello"
	var str2 = "world"
	str3 := str1 + " " + str2
	fmt.Println(str3)

	length := len(str3)
	fmt.Printf("length : %d\n", length)

	substr1 := str3[0:5]
	substr2 := str3[6:]
	fmt.Printf("substr1 : %s\n", substr1)
	fmt.Printf("substr2 : %s\n", substr2)

	str4 := reverse1(str3)
	fmt.Printf("str4 : %s\n", str4)

	str5 := reverse2(str3)
	fmt.Printf("str5 : %s\n", str5)

}

// 翻转字符串
func reverse1(str string) string {
	var result string
	length := len(str)
	for i := length - 1; i >= 0; i-- {
		//result += fmt.Sprintf("%c", str[i])
		result += string(str[i])
	}
	return result
}

// 翻转字符串
func reverse2(str string) string {
	var result []byte
	strByteArr := []byte(str)
	length := len(strByteArr)
	for i := length - 1; i >= 0; i-- {
		result = append(result, strByteArr[i])
	}
	return string(result)
}
