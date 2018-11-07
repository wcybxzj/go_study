package service

import (
	"fmt"
	"sync"
)

var (
	secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int]*Limit, 10000),
	}
)

type SecLimitMgr struct {
	//用户维度的限制
	UserLimitMap map[int]*Limit
	//IP维度的限制
	IpLimitMap map[string]*Limit
	lock       sync.Mutex
}

func antiSpam(req *SecRequest) (err error) {

	//_, ok := secKillConf.idBlackMap[req.UserId]
	//if ok {
	//	err = fmt.Errorf("invalid request")
	//	logs.Error("useId[%v] is block by id black", req.UserId)
	//	return
	//}
	//
	//_, ok = secKillConf.ipBlackMap[req.ClientAddr]
	//if ok {
	//	err = fmt.Errorf("invalid request")
	//	logs.Error("useId[%v] ip[%v] is block by ip black", req.UserId, req.ClientAddr)
	//	return
	//}

	secLimitMgr.lock.Lock()
	//uid 频率控制
	secLimit, ok := secLimitMgr.UserLimitMap[req.UserId]
	if !ok {
		secLimit = &SecLimit{}
		secLimitMgr.UserLimitMap[req.UserId] = secLimit
	}
	count := secLimit.Count(req.AccessTime.Unix())

	//if count > secKillConf.UserSecAccessLimit {
	//	err = fmt.Errorf("invalid request")
	//	return
	//}

	//secIdCount := limit.secLimit.Count(req.AccessTime.Unix())
	//minIdCount := limit.minLimit.Count(req.AccessTime.Unix())

	////ip 频率控制
	//ipLimit, ok := secLimitMgr.IpLimitMap[req.ClientAddr]
	//if !ok {
	//	ipLimit = &SecLimit{}
	//	secLimitMgr.IpLimitMap[req.ClientAddr] = ipLimit
	//}
	//
	//count = ipLimit.Count(req.AccessTime.Unix())

	secLimitMgr.lock.Unlock()
	if count > secKillConf.IPSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

	//
	//secIpCount := limit.secLimit.Count(req.AccessTime.Unix())
	//minIpCount := limit.minLimit.Count(req.AccessTime.Unix())

	//
	//if secIpCount > secKillConf.AccessLimitConf.IPSecAccessLimit {
	//	err = fmt.Errorf("invalid request")
	//	return
	//}
	//
	//if minIpCount > secKillConf.AccessLimitConf.IPMinAccessLimit {
	//	err = fmt.Errorf("invalid request")
	//	return
	//}
	//
	//if secIdCount > secKillConf.AccessLimitConf.UserSecAccessLimit {
	//	err = fmt.Errorf("invalid request")
	//	return
	//}
	//
	//if minIdCount > secKillConf.AccessLimitConf.UserMinAccessLimit {
	//	err = fmt.Errorf("invalid request")
	//	return
	//}
	return
}
