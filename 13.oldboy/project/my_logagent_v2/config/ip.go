package config

import (
	"fmt"
	"net"
)

var (
	LocalIPArray []string
)

//[192.168.1.177 172.18.0.1 172.17.0.1]
//只有第一个才是我们需要的ip
func init() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(fmt.Sprintf("get local ip failed, %v", err))
	}
	//获取本机内部ip
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				LocalIPArray = append(LocalIPArray, ipnet.IP.String())
			}
		}
	}
	//fmt.Println(LocalIPArray)
}
