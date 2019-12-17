package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"time"
)

func main() {
	//testLog()
	//testOs()
	//testTime()
	//testRegexp()
	//var i map[string]string
	//testReflect(i)
	//testUrl()
	//testMongoDB()
}

func testMongoDB() {

	dialInfo := &mgo.DialInfo{
		Addrs:          []string{"172.16.3.241:3609", "172.16.3.241:3610", "172.16.3.241:3611"},
		Username:       "username",
		Password:       "password",
		Database:       "dbname",
		ReplicaSetName: "rs",
		Mechanism:      "",
		PoolLimit:      20,
		Timeout:        10 * time.Second,
	}

	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		fmt.Println("连接MongoDB异常！", err)
	} else {
		defer session.Close()
		fmt.Println("连接MongoDB成功！")
		c := session.DB(dialInfo.Database).C("Col")

		num, _ := c.Find(bson.M{}).Count()
		fmt.Printf("共有%d个记录！", num)
	}
}

func testUrl() {
	u := "kafka://username:password@172.16.3.241:9092,172.16.3.241:9093,172.16.3.241:9094/0/dbname"
	url1, err := url.ParseRequestURI(u)
	if err != nil {
		fmt.Println("解析URL异常！", err)
	} else {
		fmt.Println("Scheme :", url1.Scheme)
		fmt.Println("Opaque :", url1.Opaque)
		fmt.Println("Username :", url1.User.Username())
		password, _ := url1.User.Password()
		fmt.Println("Password :", password)
		fmt.Println("Host :", url1.Host)
		fmt.Println("Path :", url1.Path)
		fmt.Println("RawPath :", url1.RawPath)
		fmt.Println("ForceQuery :", url1.ForceQuery)
		fmt.Println("RawQuery :", url1.RawQuery)
		fmt.Println("Fragment :", url1.Fragment)
	}

}

func testLog() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("发生错误！！", err)
		}
	}()

	//log.SetFlags(log.LstdFlags)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC | log.Lshortfile)

	log.Println("测试日志信息1")

	//log.Fatal("发生了严重的错误啦！")
	log.Panic("发生了严重的错误啦！")

	fmt.Println("严重错误之后的语句")
}

func testOs() {
	test := os.Getenv("test")
	fmt.Println(test)

	path := "/Users/caozhen/dev/go/practice/src/demo/test2/sub"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		fmt.Println("创建路径异常！", err)
	} else {
		os.Chmod(path, 0755)
		fileInfo, _ := os.Stat(path)
		fmt.Println("路径信息：")
		fmt.Println("Name：", fileInfo.Name())
		fmt.Println("IsDir：", fileInfo.IsDir())
		fmt.Println("Mode：", fileInfo.Mode())
		fmt.Println("ModTime：", fileInfo.ModTime())
		fmt.Println("Size：", fileInfo.Size())
		fmt.Println("Sys：", fileInfo.Sys())
	}

	file, err1 := os.OpenFile(path+"/1.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err1 != nil {
		fmt.Println("创建文件异常！", err1)
	} else {
		os.Chmod(path+"/1.txt", 0666)
		fmt.Println("创建文件成功！")
		str := "hello world!\n"
		num, err2 := file.Write([]byte(str))
		if err2 != nil {
			fmt.Println("将内容写入文件失败！", err)
		} else {
			fmt.Printf("将内容写入文件成功，共写入%d个字节！\n", num)
		}
	}
}

func testTime() {
	now := time.Now()
	fmt.Println("当前时间：", now)
	nextHour := now.Add(time.Hour)
	fmt.Println("一小时后", nextHour)
	nextHour = time.Date(nextHour.Year(), nextHour.Month(), nextHour.Day(), nextHour.Hour(), 0, 0, 0, nextHour.Location())
	fmt.Println("仅精确到小时", nextHour)

	nextMinute := now.Add(time.Minute)
	fmt.Println("一分钟后", nextMinute)
	nextMinute = time.Date(nextMinute.Year(), nextMinute.Month(), nextMinute.Day(), nextMinute.Hour(), nextMinute.Minute(), 0, 0, nextMinute.Location())
	timer := time.NewTimer(nextMinute.Sub(now))
	triggerTime := <-timer.C
	fmt.Println("timer触发！当前时间：", triggerTime)
	fmt.Println("timer Stop：", timer.Stop())

}

func testRegexp() {
	reg := regexp.MustCompile(`\$\{ENV_[0-9a-zA-Z_\-]+\}`)
	src := "hello, ${ENV_NAME}"
	target := reg.ReplaceAllStringFunc(src, func(s string) string {
		envName := s[6 : len(s)-1]
		envValue, find := os.LookupEnv(envName)
		if find {
			return envValue
		}

		return s
	})

	fmt.Println(target)
}

func testReflect(value interface{}) {
	v := reflect.ValueOf(value)
	if v.IsValid() {
		fmt.Println("传入的值合法！")
		kind := v.Kind()
		switch kind {
		case reflect.String:
			fmt.Println("传入的是一个字符串！值为：", value.(string))
		case reflect.Int:
			fmt.Println("传入的是一个int类型！值为：", value.(int))
		case reflect.Float32:
			fmt.Println("传入的是一个float32类型！值为：", value.(float32))
		case reflect.Float64:
			fmt.Println("传入的是一个float64类型！值为：", value.(float64))
		case reflect.Bool:
			fmt.Println("传入的是一个bool类型！值为：", value.(bool))
		default:
			fmt.Println("传入的是一个无法识别的类型！值为：", value)
		}

	} else {
		fmt.Println("传入的值非法！")
	}
}
