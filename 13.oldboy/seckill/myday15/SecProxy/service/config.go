package service

import "sync"

const(
	//商品状态正常
	ProductStatusNornal = 0
	//商品售罄
	ProductStatusSaleOut = 1
	//商品强制售罄
	ProductStatusForceSaleOut = 2
)


type RedisConf struct {
	RedisAddr        string
	RedisMaxIdle     int
	RedisMaxActive   int
	RedisIdleTimeout int
}

type EtcdConf struct {
	EtcdAddr          string
	Timeout           int
	EtcdSecKeyPrefix  string
	EtcdSecProductKey string
}

type SecSkillConf struct {
	RedisConf         RedisConf
	EtcdConf          EtcdConf
	LogPath           string
	LogLevel          string
	SecProductInfoMap map[int]*SecProductInfoConf
	RWSecProductLock  sync.RWMutex
}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int //是根据开始时间 结束时间算出来的
	Total     int
	Left      int //剩余量
}
