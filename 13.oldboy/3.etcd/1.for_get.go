package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"strconv"
	"time"
	"zuji/common/json"
	"zuji/common/timef"
)

const (
	path = "/logagent/conf/"
)

type QueueInfo struct {
	Topic 			string  `json:"Topic"`
	BrokerName      string  `json:"BrokerName"`
	QueueId         int		`json:"QueueId"`
	NextBeginOffset int64	`json:"NextBeginOffset"`
	Time            string	`json:"Time"`
}

type QueueInfoArr struct {
	QueueInfo[] QueueInfo `json:"data"`
}

var ectdConn *clientv3.Client

func initEtcd()  {
	var err error
	ectdConn, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}

	fmt.Println("connect succ")
	//defer cli.Close()
}

//benchmark
func test1()  {
	var err error
	data := "-ybx"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = ectdConn.Put(ctx, path, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	go func() {
		for {
			ctx, cancel = context.WithTimeout(context.Background(), time.Second)
			resp, err := ectdConn.Get(ctx, path)
			cancel()
			if err != nil {
				fmt.Println("get failed, err:", err)
				return
			}
			for _, ev := range resp.Kvs {
				fmt.Printf("%s : %s\n", ev.Key, ev.Value)
			}

			time.Sleep(time.Second * 1)
		}
	}()

	time.Sleep(time.Second * 100)
}

func test2()  {
	var qArr QueueInfoArr
	QueueSize := 8

	for i := 0; i < QueueSize; i++ {
		var qInfo QueueInfo
		qInfo.Topic = "Topic-Key"
		qInfo.BrokerName = "broker153"
		qInfo.QueueId = i
		qInfo.NextBeginOffset = 0
		qInfo.Time = timef.FormatTime(timef.GetTimtstamp(), "sec")
		qArr.QueueInfo = append(qArr.QueueInfo, qInfo)
	}

	//update

	SaveConfigData(qArr)
}

func SaveConfigData(qArr QueueInfoArr) {
	result, err := json.Marshal(qArr)
	if err != nil {
		fmt.Println("err:"+err.Error())
		return
	}
	data := string(result)

	_, err = ectdConn.Put(context.TODO(), path, data)
	if err != nil {
		panic("models.InitEtcd() err::" + err.Error())
		return
	}

	fmt.Println(string(result))
}

// ./etcdctl get /logagent/conf/
func test3()  {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	res, err := ectdConn.Get(ctx, path)
	cancel()
	if err != nil {
		panic("err:"+err.Error())
	}

	for _, ev := range res.Kvs {
		var qArr QueueInfoArr

		//fmt.Printf("key:%s\n", ev.Key)
		//fmt.Printf("val: %s\n", ev.Value)

		err := json.Unmarshal(ev.Value, &qArr)
		if err != nil {
			fmt.Println("err:"+err.Error())
			return
		}

		fmt.Println(qArr)
		for _, qInfo := range qArr.QueueInfo {
			fmt.Println(qInfo.Topic)
			fmt.Println(qInfo.BrokerName)
		}
	}
}

//同一个key 生成1w个不同的值
func test4()  {
	var err error
	data := "-ybx"

	for i:=0; i<10000 ; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		_, err = ectdConn.Put(ctx, path, strconv.Itoa(i) + data)
		cancel()
		if err != nil {
			fmt.Println("put failed, err:", err)
			return
		}
	}
}

//进行修改查看变化
// ./etcdctl put /logagent/conf/ 123456

// ./etcdctl get /logagent/conf/
func main() {
	initEtcd()

	//test1()

	//test2()

	//test3()

	test4()
}

