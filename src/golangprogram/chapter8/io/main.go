package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	demo1()
}

func demo1() {
	var buf bytes.Buffer
	buf.Write([]byte("Hello "))

	_, err := fmt.Fprint(&buf, "World")
	if err != nil {
		fmt.Println("fmt.Fprint failed, err :", err)
		return
	}

	buf.WriteTo(os.Stdout)
}
