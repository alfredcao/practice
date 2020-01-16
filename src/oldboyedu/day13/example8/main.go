package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"time"
)

type Tweet struct {
	User    string
	Message string
}

func main() {
	cli, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("new elasticsearch client failed, err :", err)
		return
	}

	fmt.Println("connect elasticsearch success")
	tweet := Tweet{
		User:    "alfred",
		Message: "Take Five",
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	_, err = cli.Index().
		Index("twitter").
		Type("tweet").
		Id("1").
		BodyJson(tweet).
		Do(ctx)

	if err != nil {
		fmt.Println("insert document to elasticsearch failed, err :", err)
		return
	}

	fmt.Println("insert document to elasticsearch success")
}
