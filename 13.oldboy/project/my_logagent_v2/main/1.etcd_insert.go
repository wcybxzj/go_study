package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go_study/13.oldboy/project/my_logagent_v2/config"
	"time"
)

//插入配置
func SetLogConfToEtcd() {
	//1.conf
	err := config.LoadConf("ini", config.Filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v", err)
		panic("load conf failed")
		return
	}
	//fmt.Println(config.AppConfig.EtcdKey)

	//2.etcd connect
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return
	}
	fmt.Println("connect succ")
	defer cli.Close()

	//3.logConfArr
	var logConfArr []config.CollectConf
	logConfArr = append(
		logConfArr,
		config.CollectConf{
			LogPath: "/usr/local/nginx/logs/access.log",
			Topic:   "nginx_log",
		},
	)
	logConfArr = append(
		logConfArr,
		config.CollectConf{
			LogPath: "/usr/local/nginx/logs/error.log",
			Topic:   "nginx_log_err",
		},
	)

	//4.etcd put/get
	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("json failed, ", err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, config.EtcdKey, string(data))
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, config.EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}

//对etcd进行写入,main.go才能从etcd读取到配置从而进行对设置的日志进行收集
func main() {
	//插入etcd
	SetLogConfToEtcd()
}
