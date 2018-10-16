package main

import (
	"context"
	"fmt"
	elastic "gopkg.in/olivere/elastic.v5"
)

type Tweet struct {
	User    string
	Message string
}

//插入数据后需要在elasticSearch 创建索引 twitter
func main() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://127.0.0.1:15920/"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es succ")
	for i := 0; i < 10000; i++ {
		tweet := Tweet{User: "olivere", Message: "Take Five"}
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d", i)).
			BodyJson(tweet).
			Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
			return
		}
	}

	fmt.Println("insert succ")
}
