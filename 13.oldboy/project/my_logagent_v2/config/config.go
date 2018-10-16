package config

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
)

var (
	AppConfig *Config
	Filename  = "/root/www/go_www/src/go_study/13.oldboy/project/my_logagent_v2/logagent.conf"
)

//ip必须是对外的那个
const (
	EtcdKey = "/oldboy/backend/logagent/config/192.168.1.177"
)

//总的config
type Config struct {
	LogLevel    string
	LogPath     string
	ChanSize    int
	KafkaAddr   string
	EtcdAddr    string
	EtcdKey     string
	CollectConf []CollectConf
}

//每个collect
type CollectConf struct {
	LogPath string `json:"logpath"`
	Topic   string `json:"topic"`
}

//只加载collect部分conf
func loadCollectConf(conf config.Configer) (err error) {
	var cc CollectConf
	cc.LogPath = conf.String("collect::log_path")
	if len(cc.LogPath) == 0 {
		err = errors.New("invalid collect::log_path")
		return
	}

	cc.Topic = conf.String("collect::topic")
	if len(cc.Topic) == 0 {
		err = errors.New("invalid collect::topic")
		return
	}

	AppConfig.CollectConf = append(AppConfig.CollectConf, cc)
	return
}

//加载全部conf
func LoadConf(confType, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	AppConfig = &Config{}
	AppConfig.LogLevel = conf.String("logs::log_level")
	if len(AppConfig.LogLevel) == 0 {
		AppConfig.LogLevel = "debug"
	}

	AppConfig.LogPath = conf.String("logs::log_path")
	if len(AppConfig.LogPath) == 0 {
		AppConfig.LogPath = "/tmp/logs"
	}

	AppConfig.ChanSize, err = conf.Int("logs::chan_size")
	if err != nil {
		AppConfig.ChanSize = 100
	}

	AppConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(AppConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	AppConfig.EtcdAddr = conf.String("etcd::server_addr")
	if len(AppConfig.EtcdAddr) == 0 {
		err = fmt.Errorf("invalid etcd addr")
		return
	}

	AppConfig.EtcdKey = conf.String("etcd::config_key")
	if err != nil {
		err = fmt.Errorf("invalid etcd config_key")
		return
	}

	//err = loadCollectConf(conf)
	//if err != nil {
	//	fmt.Printf("load collect conf failed,err:%v", err)
	//	return
	//}
	return
}
