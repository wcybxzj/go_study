package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go_study/13.oldboy/project/my_logagent_v2/config"
	"time"
)

func DelLogConfToEtcd() {
	//1.conf
	err := config.LoadConf("ini", config.Filename)
	if err != nil {
		fmt.Printf("load conf failed, err:%v", err)
		panic("load conf failed")
		return
	}

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Delete(ctx, config.EtcdKey)
	cancel()

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
	//删除etcd, 用来测试main中watch, 发现etcd配置修改
	DelLogConfToEtcd()
}
