package main

import (
	"encoding/json"
	"fmt"
)


type PhpResponse struct {
	//Data interface{} `json:"data" binding:"required"`
	Code interface{} `json:"code" binding:"required"`
	//Msg string `json:"msg" binding:"required"`
}

func test(str string)  {
	fmt.Println("start func")
	END:
	for i:=0; i<3; i++ {

		fmt.Println("i:",string(i))
		var phpResp PhpResponse
		strByte := []byte(str)

		err := json.Unmarshal(strByte, &phpResp)
		if err!=nil	{
			fmt.Println(err.Error())
		}

		switch phpResp.Code.(type) {
			case float64:
				fmt.Println("flloat64:")
				code := int((phpResp.Code).(float64) )
				if code == 0 {
					fmt.Println("ok")
					break END
				}
			case string:
				fmt.Println("string:")
				if phpResp.Code == "0" {
					fmt.Println("ok")
					break END
				}
		}
	}
	fmt.Println("function is done!")
}

/*
"data":"{"order_no":"SOBB0452597895633","order_event":"order_deliveried"}",
"key_type":"nqy_order_deliveried",
"md5":"b8b352efd8adad9a14064bd23b05b3d9"
*/
func main() {

str1:= `{"code":0,
"msg":"支付宝订单信息同步成功",
"data":{"name":"ybx"}
}`

str2:= `{"code":1,
"msg":"支付宝订单信息同步成功",
"data":{"name":"ybx"}
}`

str3:= `{"code":"0",
"msg":"支付宝订单信息同步成功",
"data":{"name":"ybx"}
}`

str4:= `{"code":"1",
"msg":"支付宝订单信息同步成功",
"data":{"name":"ybx"}
}`



	test(str1)

	test(str2)

	test(str3)

	test(str4)
}
