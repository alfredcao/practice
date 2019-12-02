package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	sub := Sub{}
	sub.Foo()
	sub.Bar()

	job := Job{"xxx", log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)}
	job.Printf("1111")
}

type Base struct {
	Name string
}

func (base Base) Foo() {
	fmt.Println("foo in base")
}

func (base Base) Bar() {
	fmt.Println("bar in base")
}

type Sub struct {
	Base
	Address string
}

func (sub Sub) Foo() {
	fmt.Println("foo in sub")
	fmt.Println(sub.Address)
}

type Job struct {
	Command string
	*log.Logger
}

func (job *Job) Start() {

}
