package config

import (
	"fmt"

	"github.com/astaxie/beego/config"
)

type LogConfigSt struct {
	KafkaAddr  string
	ESAddr     string
	LogPath    string
	LogLevel   string
	KafkaTopic string
}

var (
	LogConfig *LogConfigSt
	Filename  = "/root/www/go_www/src/go_study/13.oldboy/project/log_transfer/log_transfer.conf"
)

func InitConfig(confType string, filename string) (err error) {
	conf, err := config.NewConfig(confType, filename)
	if err != nil {
		fmt.Println("new config failed, err:", err)
		return
	}

	LogConfig = &LogConfigSt{}
	LogConfig.LogLevel = conf.String("logs::log_level")
	if len(LogConfig.LogLevel) == 0 {
		LogConfig.LogLevel = "debug"
	}

	LogConfig.LogPath = conf.String("logs::log_path")
	if len(LogConfig.LogPath) == 0 {
		LogConfig.LogPath = "./logs"
	}

	LogConfig.KafkaAddr = conf.String("kafka::server_addr")
	if len(LogConfig.KafkaAddr) == 0 {
		err = fmt.Errorf("invalid kafka addr")
		return
	}

	LogConfig.KafkaTopic = conf.String("kafka::topic")
	if len(LogConfig.KafkaTopic) == 0 {
		err = fmt.Errorf("invalid kafka topic")
		return
	}

	LogConfig.ESAddr = conf.String("es::server_addr")
	if len(LogConfig.ESAddr) == 0 {
		err = fmt.Errorf("invalid es addr")
		return
	}

	return
}
