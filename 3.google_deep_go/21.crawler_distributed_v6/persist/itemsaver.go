package persist

import (
	"context"
	"errors"
	"go_study/3.google_deep_go/21.crawler_distributed_v6/engine"
	"gopkg.in/olivere/elastic.v5"
)

//存入elasticseatch 可以用http.Post
//或者
//用elasticseatch go client 非官方用户的版本
//go get gopkg.in/olivere/elastic.v5
func Save(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" {
		return errors.New("must supply Type")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	//创建或者修改数据
	//index: 相当于数据库		   ---> 开发人员用配置进行执行
	//type:	相当于表 id:相当于id ---> 是parser给出
	_, err := indexService.Do(context.Background())

	if err != nil {
		return err
	}

	//打印结构体时候架上字段名
	//fmt.Printf("%+v", resp)

	return nil
}
