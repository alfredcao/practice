package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strings"
)

func main() {
	//processXml()
	processXml2()
}

func processXml() {
	post := Post{}
	err := xml.Unmarshal([]byte(postXml), &post)
	if err != nil {
		fmt.Println("xml unmarshal failed, err :", err)
		return
	}

	fmt.Println(post)

	post.RowXml = ""
	xmlBytes, err := xml.Marshal(post)
	if err != nil {
		fmt.Println("xml marshal failed, err :", err)
		return
	}

	fmt.Println("marshal result :")
	fmt.Println(string(xmlBytes))
}

/**
流式处理
*/
func processXml2() {
	decoder := xml.NewDecoder(strings.NewReader(postXml))
	var comments []Comment
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("decode failed, err :", err)
			return
		}

		switch se := token.(type) {
		case xml.StartElement:
			elementName := se.Name.Local
			if elementName == "comment" {
				var comment Comment
				err = decoder.DecodeElement(&comment, &se)
				if err != nil {
					fmt.Println("decode element failed, err :", err)
					return
				}
				comments = append(comments, comment)
			}
		}
	}
	fmt.Println(comments)

	buf := bytes.NewBufferString("")
	encoder := xml.NewEncoder(buf)
	err := encoder.Encode(comments)

	if err != nil {
		fmt.Println("encode failed, err :", err)
		return
	}

	fmt.Println(buf.String())

}

const postXml = `
<?xml version="1.0" encoding="utf-8"?>
<post id="123">
	<content>hello</content>
	<author id="1">leo</author>
	<comments>
		<comment>
			<content>你好</content>
			<author id="2">zhangsan</author>
		</comment>
		<comment>
			<content>aliaduoguozayimasi</content>
			<author id="3">lisi</author>
		</comment>
	</comments>
</post>
`

type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       int       `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
	RowXml   string    `xml:",innerxml"`
}

type Author struct {
	XMLName xml.Name `xml:"author"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:",chardata"`
}

type Comment struct {
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}
