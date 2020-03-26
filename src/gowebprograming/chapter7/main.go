package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	processXml()

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
