package persist

import (
	"context"
	"elastic-release-branch.v5"
	"encoding/json"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/engine"
	"go_study/3.google_deep_go/19.crawler_concurrent_itemsaver_v5/model"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Name:       "安静的雪",
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Car:        "未购车",
			Education:  "大学本科",
			Hokou:      "山东菏泽",
			Marriage:   "离异",
			House:      "已购房",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	const index = "dating_test"

	//保存数据
	err = Save(client, index, expected)
	if err != nil {
		panic(err)
	}

	//获取数据
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	//t.Logf("%s", *resp.Source)

	var actual engine.Item
	json.Unmarshal([]byte(*resp.Source), &actual)
	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	//t.Logf("%+v", actual)
	//t.Logf("%+v", expected)

	if actual != expected {
		t.Errorf("got %v; expected:%v", actual, expected)
	}
}
