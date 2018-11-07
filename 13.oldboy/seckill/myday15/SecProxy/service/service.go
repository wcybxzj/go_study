package service

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/garyburd/redigo/redis"
	"github.com/prometheus/common/log"
	"strconv"
	"time"
)

var (
	secKillConf *SecSkillConf
)

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	loadIdBlackList()
	logs.Debug("init service succ, config:%v", secKillConf)

	//initProxy2LayerRedis()

	secKillConf.secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int]*s, 10000),
		IpLimitMap:   make(map[string]*Limit, 10000),
	}
}

func initBlackRedis() (err error) {
	secKillConf.blackRedisPool = &redis.Pool{
		MaxIdle:     secKillConf.RedisBlackConf.RedisMaxIdle,
		MaxActive:   secKillConf.RedisBlackConf.RedisMaxActive,
		IdleTimeout: time.Duration(secKillConf.RedisBlackConf.RedisIdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", secKillConf.RedisBlackConf.RedisAddr)
		},
	}

	conn := secKillConf.blackRedisPool.Get()
	defer conn.Close()

	_, err = conn.Do("ping")
	if err != nil {
		logs.Error("ping redis failed, err:%v", err)
		return
	}

	return
}

//加载黑名单(全量)
func loadIdBlackList() (err error) {
	err = initBlackRedis()
	if err != nil {
		log.Error("init black redis, failed, err:%v", err)
		return
	}

	conn := secKillConf.blackRedisPool.Get()
	defer conn.Close()

	//id
	reply, err := conn.Do("hgetall", "idblacklist")
	idlist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}

	for _, v := range idlist {
		id, err := strconv.Atoi(v)
		if err != nil {
			log.Warn("invalid user id[%v]", id)
			continue
		}
		secKillConf.idBlackMap[id] = true
	}

	//ip
	reply, err = conn.Do("hgetall", "ipblacklist")
	iplist, err := redis.Strings(reply, err)
	if err != nil {
		logs.Warn("hget all failed, err:%v", err)
		return
	}

	for _, v := range iplist {
		secKillConf.ipBlackMap[v] = true
	}

	go SyncIpBlackList()
	go SyncIdBlackList()
	return
}

//同步黑名单数据(redis阻塞获取+批量更新+增量方式)
//累计获取100个新的黑名单  或者 大于5秒 就会强制更新
//定时将redis中获取的黑名单同步到服务
func SyncIpBlackList() {
	var ipList []string
	lastTime := time.Now().Unix()

	for {
		conn := secKillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackiplist", time.Second)
		ip, err := redis.String(reply, err)
		if err != nil {
			continue
		}

		curTime := time.Now().Unix()
		ipList = append(ipList, ip)

		if len(ipList) > 100 || curTime-lastTime > 5 {
			secKillConf.RWBlackLock.Lock()
			for _, v := range ipList {
				secKillConf.ipBlackMap[v] = true
			}
			secKillConf.RWBlackLock.Unlock()
			lastTime = curTime
			logs.Info("sync ip list from redis succ, ip[%v]", ipList)
		}
	}
}

//和SyncIpBlackList 一样的策略 不写了
func SyncIdBlackList() {
	for {
		conn := secKillConf.blackRedisPool.Get()
		defer conn.Close()
		reply, err := conn.Do("BLPOP", "blackidlist", time.Second)
		id, err := redis.Int(reply, err)
		if err != nil {
			continue
		}
		secKillConf.RWBlackLock.Lock()
		secKillConf.idBlackMap[id] = true
		secKillConf.RWBlackLock.Unlock()
		logs.Info("sync id list from redis succ, id:%v", id)

	}
}

func SecInfoList() (data []map[string]interface{}, code int, err error) {
	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	for _, v := range secKillConf.SecProductInfoMap {
		item, _, err := SecInfoById(v.ProductId)
		if err != nil {
			logs.Error("get product_id[%d] failed, err:%v", v.ProductId, err)
			continue
		}

		logs.Debug("get product[%d]， result[%v], all[%v] v[%v]", v.ProductId, item, secKillConf.SecProductInfoMap, v)
		data = append(data, item)
	}

	return
}

func SecInfo(productId int) (data []map[string]interface{}, code int, err error) {
	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	item, code, err := SecInfoById(productId)
	if err != nil {
		return
	}

	data = append(data, item)

	return
}

func SecInfoById(productId int) (data map[string]interface{}, code int, err error) {
	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	v, ok := secKillConf.SecProductInfoMap[productId]
	if !ok {
		code = ErrNotFoundProductId
		err = fmt.Errorf("not found product_id:%d", productId)
		return
	}

	start := false
	end := false
	status := "success"

	now := time.Now().Unix()

	if now-v.StartTime < 0 {
		start = false
		end = false
		status = "sec kill is not start"
		code = ErrActiveNotStart
	}

	if now-v.StartTime > 0 {
		start = true
		end = false
	}

	if now-v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
		code = ErrActiveAlreadyEnd
	}

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut {
		start = false
		end = true
		status = "Product is sale out"
		code = ErrActiveSaleOut
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = start
	data["end_time"] = end
	data["status"] = status

	return
}

func userCheck(req *SecRequest) (err error) {

	//found := false
	//for _, refer := range secKillConf.ReferWhiteList {
	//	if refer == req.ClientRefence {
	//		found = true
	//		break
	//	}
	//}
	//
	//if !found {
	//	err = fmt.Errorf("invalid request")
	//	logs.Warn("user[%d] is reject by refer, req[%v]", req.UserId, req)
	//	return
	//}
	//
	authData := fmt.Sprintf("%d:%s", req.UserId, secKillConf.CookieSecretKey)
	authSign := fmt.Sprintf("%x", md5.Sum([]byte(authData)))

	if authSign != req.UserAuthSign {
		err = fmt.Errorf("invalid user cookie auth")
		return
	}
	return
}

func SecKill(req *SecRequest) (data map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	//err = userCheck(req)
	//if err != nil {
	//	code = ErrUserCheckAuthFailed
	//	logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	//	return
	//}

	err = antiSpam(req)
	if err != nil {
		code = ErrUserServiceBusy
		logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
		return
	}

	data, code, err = SecInfoById(req.ProductId)
	if err != nil {
		logs.Warn("userId[%d] secInfoBy Id failed, req[%v]", req.UserId, req)
		return
	}

	if code != 0 {
		logs.Warn("userId[%d] secInfoByid failed, code[%d] req[%v]", req.UserId, code, req)
		return
	}

	//userKey := fmt.Sprintf("%s_%s", req.UserId, req.ProductId)
	//secKillConf.UserConnMap[userKey] = req.ResultChan
	//
	//secKillConf.SecReqChan <- req
	//
	//ticker := time.NewTicker(time.Second * 10)
	//
	//defer func() {
	//	ticker.Stop()
	//	secKillConf.UserConnMapLock.Lock()
	//	delete(secKillConf.UserConnMap, userKey)
	//	secKillConf.UserConnMapLock.Unlock()
	//}()
	//
	//select {
	//case <-ticker.C:
	//	code = ErrProcessTimeout
	//	err = fmt.Errorf("request timeout")
	//
	//	return
	//case <-req.CloseNotify:
	//	code = ErrClientClosed
	//	err = fmt.Errorf("client already closed")
	//	return
	//case result := <-req.ResultChan:
	//	code = result.Code
	//	data["product_id"] = result.ProductId
	//	data["token"] = result.Token
	//	data["user_id"] = result.UserId
	//
	//	return
	//}
	//
	return
}
