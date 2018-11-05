package service

import (
	"fmt"
	"sync"
)

var (
	secLimitMgr = &SecLimitMgr{
		UserLimitMap: make(map[int]*SecLimit, 10000),
	}
)

type SecLimitMgr struct {
	UserLimitMap map[int]*SecLimit
	lock         sync.Mutex
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

	//secIdCount := limit.secLimit.Count(req.AccessTime.Unix())
	//minIdCount := limit.minLimit.Count(req.AccessTime.Unix())
	//
	////ip 频率控制
	//limit, ok = secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr]
	//if !ok {
	//	limit = &Limit{
	//		secLimit: &SecLimit{},
	//		minLimit: &MinLimit{},
	//	}
	//	secKillConf.secLimitMgr.IpLimitMap[req.ClientAddr] = limit
	//}
	//
	//secIpCount := limit.secLimit.Count(req.AccessTime.Unix())
	//minIpCount := limit.minLimit.Count(req.AccessTime.Unix())

	count := secLimit.Count(req.AccessTime.Unix())

	secLimitMgr.lock.Unlock()

	if count > secKillConf.UserSecAccessLimit {
		err = fmt.Errorf("invalid request")
		return
	}

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
