package main

import (
	"bufio"
	"flag"
	"fmt"
	"golangprogram/chapter2/sorter/algorithm/bubblesort"
	"golangprogram/chapter2/sorter/algorithm/qsort"
	"io"
	"os"
	"strconv"
)

var inFile *string = flag.String("i", "inFile", "File contains values For sorting")
var outFile *string = flag.String("o", "outFile", "File to receive sorted values")
var algorithm *string = flag.String("a", "algorithm", "Sort algorithm")

func main() {
	flag.Parse()
	fmt.Println("inFile =", *inFile, "outFile =", *outFile, "algorithm =", *algorithm)
	values, err := readFile(*inFile)
	if err != nil {
		fmt.Println("读取输入文件异常！", err)
		return
	} else {
		fmt.Println("读取到输入文件内容！", values)
	}

	switch *algorithm {
	case "qsort":
		qsort.QuickSort(values)
	case "bubblesort":
		bubblesort.BubbleSort(values)
	default:
		fmt.Println("不支持的排序算法！")
		return
	}

	err = writeFile(values, *outFile)
	if err != nil {
		fmt.Println("写入输出文件异常！", err)
		return
	} else {
		fmt.Println("写入输出文件成功！", values)
	}
}

func readFile(filePath string) (values []int, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed to open the input file ", filePath)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	values = make([]int, 0)
	for {
		line, isPrefix, err1 := br.ReadLine()
		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, seems unexpected.")
			return
		}

		str := string(line)
		num, err2 := strconv.Atoi(str)
		if err2 != nil {
			err = err2
			return
		}
		values = append(values, num)
	}
	return
}

func writeFile(values []int, filePath string) (err error) {
	file, err1 := os.Create(filePath)
	if err1 != nil {
		err = err1
		return
	}

	defer file.Close()

	for _, value := range values {
		_, err2 := file.WriteString(strconv.Itoa(value) + "\n")
		if err2 != nil {
			err = err2
			return
		}
	}

	return nil
}
