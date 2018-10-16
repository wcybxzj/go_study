package persist

import (
	"context"
	"errors"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/engine"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

//itemsaver是一个goroutine
func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false)) //需要集群状态的维护,关闭sniff
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
			err := Save(client, index, item)
			if err != nil {
				//保存失败会报错
				log.Printf("item saver error"+
					"saving item %d: %v", itemCount, item)
			}
		}
	}()
	return out, nil
}

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
