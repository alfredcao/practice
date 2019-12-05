package main

import "fmt"

func main() {
	//mulTable() // 九九乘法表
	//completeNum() // 完数
	//palindrome("上海上") // 回文
	count("sfj上华菱精工   439587  sdf")
}

func mulTable() {
	for i := 1; i <= 9; i++ {
		for j := i; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println("")
	}
}

func completeNum() {
	for i := 1; i <= 1000; i++ {
		sum := 0
		for j := 1; j < i; j++ {
			if i%j == 0 {
				sum += j
			}
		}

		if sum == i {
			fmt.Printf("%d\t", i)
		}
	}
}

func palindrome(text string) {
	mytext := []rune(text)
	length := len(mytext)
	flag := true
	for i := 0; i < length; i++ {
		if mytext[i] != mytext[length-1-i] {
			flag = false
			break
		}
	}

	if flag {
		fmt.Printf("%s 是 回文\n", text)
	} else {
		fmt.Printf("%s 非 回文\n", text)
	}

}

func count(text string) {
	englishCount := 0
	spaceCount := 0
	digitalCount := 0
	otherCount := 0
	mytext := []rune(text)
	length := len(mytext)
	for i := 0; i < length; i++ {
		x := mytext[i]
		switch {
		case (x >= 65 && x <= 90) || (x >= 97 && x <= 122):
			englishCount += 1
		case x == 32:
			spaceCount += 1
		case x >= 48 && x <= 57:
			digitalCount += 1
		default:
			otherCount += 1
		}
	}

	fmt.Printf("英文字符个数 -> %d，空格个数 -> %d，数字个数 -> %d，其它字符个数 -> %d\n", englishCount, spaceCount, digitalCount, otherCount)
}
