package service

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

const (
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
	RedisBlackConf       RedisConf
	RedisProxy2LayerConf RedisConf
	RedisLayer2ProxyConf RedisConf

	RedisConf          RedisConf
	EtcdConf           EtcdConf
	LogPath            string
	LogLevel           string
	SecProductInfoMap  map[int]*SecProductInfoConf
	RWSecProductLock   sync.RWMutex
	CookieSecretKey    string
	UserSecAccessLimit int //用户每秒访问频率
	ReferWhiteList     []string
	IPSecAccessLimit   int

	//AccessLimitConf      AccessLimitConf

	blackRedisPool       *redis.Pool //黑名单redis实例
	proxy2LayerRedisPool *redis.Pool //接入层到逻辑层redis实例
	layer2ProxyRedisPool *redis.Pool //逻辑层到接入层redis实例

	secLimitMgr *SecLimitMgr

	//黑名单
	ipBlackMap map[string]bool
	idBlackMap map[int]bool
	//操作黑名单map的锁
	RWBlackLock sync.RWMutex

	//WriteProxy2LayerGoroutineNum int
	//ReadProxy2LayerGoroutineNum  int
	//
	//SecReqChan     chan *SecRequest
	//SecReqChanSize int
	//
	//UserConnMap     map[string]chan *SecResult
	//UserConnMapLock sync.Mutex

}

type SecProductInfoConf struct {
	ProductId int
	StartTime int64
	EndTime   int64
	Status    int //是根据开始时间 结束时间算出来的
	Total     int
	Left      int //剩余量
}

type SecRequest struct {
	ProductId    int
	Source       string //来源
	AuthCode     string //
	SecTime      string //当前时间
	Nance        string
	UserId       int
	UserAuthSign string //登录后的cookie
	AccessTime   time.Time

	ClientAddr    string
	ClientRefence string
	CloseNotify   <-chan bool

	//ResultChan chan *SecResult
}
