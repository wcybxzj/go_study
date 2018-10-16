package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	etcd_client "go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"go_study/13.oldboy/project/my_logagent_v2/config"
	"go_study/13.oldboy/project/my_logagent_v2/tailf"
	"strings"
	"time"
)

type EtcdClient struct {
	client *etcd_client.Client
	keys   []string
}

var (
	etcdClient *EtcdClient
)

func InitEtcd(addr string, key string) (collectConf []config.CollectConf, err error) {
	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	//fmt.Println("connect succ")

	etcdClient = &EtcdClient{
		client: cli,
	}

	if strings.HasSuffix(key, "/") == false {
		key = key + "/"
	}

	//每个etcdKey中的value可以有多个配置
	//每个etcdKey中的的Key是
	//循环每一个需要监控的etcdKey
	for _, ip := range config.LocalIPArray {
		etcdKey := fmt.Sprintf("%s%s", key, ip)
		etcdClient.keys = append(etcdClient.keys, etcdKey)
		//fmt.Println(etcKey)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		resp, err := cli.Get(ctx, etcdKey)
		if err != nil {
			logs.Error("client get from etcd failed, err:%v", err)
			panic(err)
			continue
		}
		cancel()
		logs.Debug("resp from etcd:%v", resp.Kvs)
		for _, v := range resp.Kvs {
			//fmt.Println(k, v)
			if string(v.Key) == etcdKey {
				json.Unmarshal(v.Value, &collectConf)
				if err != nil {
					logs.Error("Unmarshal failed, err:%v", err)
					continue
				}
				logs.Debug("log cond is %v", collectConf)
			}
		}
	}

	InitEtcdWathcher()
	return
}

//启动多个go来监控多个配置节点是否变化
func InitEtcdWathcher() {
	//fmt.Println("InitEtcdWathcher")
	logs.Debug("InitEtcdWathcher")
	//fmt.Println("etcdClient.keys", etcdClient.keys)
	logs.Debug("etcdClient.keys", etcdClient.keys)

	for _, key := range etcdClient.keys {
		//fmt.Println(key)
		go watchKey(key)
	}
}

//监控etcd key的变化
func watchKey(key string) {
	fmt.Println("watched key:", key)
	logs.Debug("watched key:", key)

	cli, err := etcd_client.New(etcd_client.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		logs.Error("connect failed, err:", err)
		return
	}

	//通过etcd watch来获取最新的配置
	for {
		rch := cli.Watch(context.Background(), key)
		var collectConf []config.CollectConf
		var getConfSucc = true
		_ = getConfSucc
		for wresp := range rch {
			for _, ev := range wresp.Events {
				//删除
				if int32(ev.Type) == int32(mvccpb.DELETE) {
					logs.Warn("key[%s] 's config deleted", key)
					continue
				}
				//更新
				if int32(ev.Type) == int32(mvccpb.PUT) && string(ev.Kv.Key) == key {
					err = json.Unmarshal(ev.Kv.Value, &collectConf)
					if err != nil {
						logs.Warn("key[%s] 's config deleted", key)
						getConfSucc = false
						continue
					}
				}
				logs.Debug("action:%s key:%q : value:%q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			}

			if getConfSucc {
				logs.Debug("get config from etcd succ", collectConf)
				tailf.UpdateConfig(collectConf)
			}
		}
	}
}
