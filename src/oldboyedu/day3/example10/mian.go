package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//swithSimpleSample()
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100)
	flag := false
	for {
		var num int
		fmt.Scanf("%d\n", &num)
		switch {
		case randomNum == num:
			fmt.Println("猜对啦！")
			flag = true
		case num < randomNum:
			fmt.Println("猜小了！")
		case num > randomNum:
			fmt.Println("猜大了！")
		}

		if flag {
			break
		}
	}

}

func swithSimpleSample() {
	i := 0
	switch i {
	case 0:
		fallthrough
	case 10:
		fmt.Println("hello!")
	case 20:
		fmt.Println("I fine!")
	default:
		fmt.Println("goodbye!")
	}
}
