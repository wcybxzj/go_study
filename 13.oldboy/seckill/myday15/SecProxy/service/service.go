package service

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"time"
)

var (
	secKillConf *SecSkillConf
)

func InitService(serviceConf *SecSkillConf) {
	secKillConf = serviceConf
	logs.Debug("init service succ, config:%v", secKillConf)
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
	}

	if now-v.StartTime > 0 {
		start = true
		end = false
	}

	if now-v.EndTime > 0 {
		start = false
		end = true
		status = "sec kill is already end"
	}

	if v.Status == ProductStatusForceSaleOut || v.Status == ProductStatusSaleOut {
		start = false
		end = true
		status = "Product is sale out"
	}

	data = make(map[string]interface{})
	data["product_id"] = productId
	data["start_time"] = start
	data["end_time"] = end
	data["status"] = status

	return
}

func SecKill(product_id int, source, authcode, sec_time, nance string) (data map[string]interface{}, code int, err error) {

	secKillConf.RWSecProductLock.RLock()
	defer secKillConf.RWSecProductLock.RUnlock()

	//err = userCheck(req)
	//if err != nil {
	//	code = ErrUserCheckAuthFailed
	//	logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	//	return
	//}
	//
	//err = antiSpam(req)
	//if err != nil {
	//	code = ErrUserServiceBusy
	//	logs.Warn("userId[%d] invalid, check failed, req[%v]", req.UserId, req)
	//	return
	//}
	//
	//data, code, err = SecInfoById(req.ProductId)
	//if err != nil {
	//	logs.Warn("userId[%d] secInfoBy Id failed, req[%v]", req.UserId, req)
	//	return
	//}
	//
	//if code != 0 {
	//	logs.Warn("userId[%d] secInfoByid failed, code[%d] req[%v]", req.UserId, code, req)
	//	return
	//}
	//
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
