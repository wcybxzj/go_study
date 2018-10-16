package main

import (
	"fmt"
	"time"
)

const (
	Female = 1
	Male   = 2
)

/*
1537956109

2018
September
26
1

2018/09/26 18/01/49
*/
func test1() {

	now := time.Now()

	//获取当前时间
	//1970到现在的秒数
	fmt.Println(now.Unix())

	fmt.Println()

	//Day, Minute, Month, Year
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Minute())

	fmt.Println()

	fmt.Printf("%02d/%02d/%02d %02d/%02d/%02d",
		now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

}

/*
秒<--*1000--微秒<--*1000--毫秒<--*1000--纳秒
单位是纳秒
time.Duration
*/
func test2() {
	now := time.Now()
	fmt.Println(now.Format("02/1/2006 15:04"))
	fmt.Println(now.Format("2006/1/02 15:04"))
	fmt.Println(now.Format("2006/1/02"))
}

func main() {
	//test1()
	test2()
}
