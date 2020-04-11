package main

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"time"
)

func main() {
	c, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		fmt.Println("new cache failed, err :", err)
		return
	}

	err = c.Put("key1", "value1", 10*time.Second)
	if err != nil {
		fmt.Println("put cache failed, err :", err)
		return
	}

	v := c.Get("key1")
	fmt.Println("value :", v)
	//time.Sleep(11 * time.Second)
	fmt.Println(c.IsExist("key1"))
	c.Delete("key1")
	fmt.Println(c.IsExist("key1"))

}
