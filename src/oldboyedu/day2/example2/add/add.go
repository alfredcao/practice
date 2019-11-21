package add

import (
	"fmt"
	_ "oldboyedu/day2/example2/sub"
)

var Name = "xxxxx"
var Age = 10

func init() {
	fmt.Println("add init")
	Name = "hello world"
	Age = 100
}
