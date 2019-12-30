package main

import (
	"flag"
	"fmt"
)

func main() {
	var confPath string
	var logLevel int
	flag.StringVar(&confPath, "c", "var/conf/config.xml", "please input config file path!")
	flag.IntVar(&logLevel, "l", 10, "please input log level!")

	flag.Parse()
	fmt.Println(confPath)
	fmt.Println(logLevel)
}
