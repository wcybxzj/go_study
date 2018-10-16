package log

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"go_study/13.oldboy/project/my_logagent_v1/config"
)

func convertLogLevel(level string) int {
	switch level {
	case "debug":
		return logs.LevelDebug
	case "warn":
		return logs.LevelWarn
	case "info":
		return logs.LevelInfo
	case "trace":
		return logs.LevelTrace
	}
	return logs.LevelDebug
}

//"/root/www/go_www/src/go_study/13.oldboy/day11/logs/1.log"
func InitLogger() (err error) {
	c := make(map[string]interface{})
	c["filename"] = config.AppConfig.LogPath
	c["level"] = convertLogLevel(config.AppConfig.LogLevel)

	configStr, err := json.Marshal(c)
	if err != nil {
		fmt.Println("marshal failed, err:", err)
		return
	}

	logs.SetLogger(logs.AdapterFile, string(configStr))
	return
}
