package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("/tmp/1.txt")
	if err != nil {
		fmt.Println("打开文件错误！", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var isLastLine = false
	for {
		var alphaCount, numCount, blankCount, otherCount = 0, 0, 0, 0
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				isLastLine = true
			} else {
				fmt.Println("读取文件错误！", err)
				return
			}
		}

		charArr := []rune(line)
		for _, char := range charArr {
			switch {
			case char >= 'a' && char <= 'z':
				fallthrough
			case char >= 'A' && char <= 'Z':
				alphaCount++
			case char >= '0' && char <= '9':
				numCount++
			case char == ' ':
				fallthrough
			case char == '\t':
				blankCount++
			default:
				otherCount++
			}
		}

		fmt.Printf("linecontent : %s\nalpha : %d, number : %d, blank : %d, other : %d\n", line, alphaCount, numCount, blankCount, otherCount)
		if isLastLine {
			break
		}
	}
}
