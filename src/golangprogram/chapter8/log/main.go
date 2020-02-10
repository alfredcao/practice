package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace *log.Logger
	Info  *log.Logger
	Warn  *log.Logger
	Error *log.Logger
)

func init() {
	log.SetPrefix("日志前缀：")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	file, err := os.OpenFile("./var/log/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("open error file failed, err :", err)
	}

	Trace = log.New(ioutil.Discard, "TRACE : ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(os.Stdout, "INFO : ", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stdout, "WARN : ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	//demo1()
	demo2()
}

func demo1() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("err :", err)
		}
	}()
	log.Println("normal message")
	//log.Fatalln("fatal message")
	log.Panicln("panic message")
}

func demo2() {
	Trace.Println("trace message")
	Info.Println("info message")
	Warn.Println("warn message")
	Error.Println("error message")
}
