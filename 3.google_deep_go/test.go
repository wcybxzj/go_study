package main

import (
	"fmt"
	"unicode"
	"zuji/common/json"
)

//判断是否为汉字
func test1()  {
	s := "Hello 世界！"
	for _, r := range s {
		// 判断字符是否为汉字
		if unicode.Is(unicode.Scripts["Han"], r) {
			fmt.Printf("%c", r) // 世界
		}
	}
}

//slice
func test2()  {
	arr := [...]int{100, 1, 2, 3, 4, 5, 6, 7}
	var i int
	var v int
	for i, v = range arr  {
		fmt.Printf("i:%d v:%d",i, v)
		break
	}
	fmt.Printf("==============")
	fmt.Printf("i:%d v:%d",i, v)
}

func test3_1(data[][]int)  {
	data[0][0]=123
}

func test3()  {
	data := make([][]int, 10)
	for i, _ := range data {
		data[i] = make([]int, 2)
	}

	test3_1(data)

	fmt.Println(data)
}


type MoblieUid struct{
	Moblie int
	UserId int
}

//struct slice
func test4() (s1 []MoblieUid){
	s := make([]MoblieUid, 0)

	var tmp MoblieUid
	tmp.Moblie = 123

	s = append(s, tmp)

	tmp.Moblie = 789
	tmp.UserId = 110
	s = append(s, tmp)

	tmp.Moblie = 100
	tmp.UserId = 101
	s = append(s, tmp)

	//fmt.Println(s)
	return s
}

//slice
func test5()  {

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)


}

//Date
func test6()  {
}

func test7() {
	data := make([][]string, 0)
	data = make([][]string, 2)

	data[0] = make([]string, 2)
	data[1] = make([]string, 2)

	data[0][0]="a"
	data[0][1]=""
	data[1][0]="c"
	data[1][1]=""
	//fmt.Println(data)

	for _,v :=range data{
		fmt.Println(v[0])

		if v[1]=="" {
			fmt.Println("v[1] is empty")
		}
	}

	/*
	excelData[0][0]="a"
	excelData[0][1]="b"
	excelData[1][0]="c"
	excelData[1][1]="d"

	fmt.Println(excelData)
	*/
}

func test8() (rowNum int) {
	return rowNum
}

func test9()  {
	var s1 string
	data := make(map[string]interface{})
	data["name"]="ybx123"
	s1 = (data["name"]).(string)
	fmt.Println(s1)
}


type OrderResponse struct {
	Data interface{} `json:"data"`
	Code string `json:"code"`
	Msg string `json:"msg"`
}

func test10() {
	var str string
	str = `{ "data": [], "code": "0", "msg": "成功" }`
	res := OrderResponse{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res.Msg)
	fmt.Println(res.Code)
	fmt.Println(res.Data)
}

func test11()  {
	//checkedExcelData := make([][]string, 0)
	var checkedExcelData [][]string

	ss1 := make([]string,0)
	ss1 = append(ss1, "1111")
	ss1 = append(ss1, "22222")
	ss1 = append(ss1, "3333")

	ss2 := make([]string,0)
	ss2 = append(ss2, "444")
	ss2 = 	append(ss2, "555")
	ss2 = append(ss2, "6666")

	checkedExcelData = append(checkedExcelData, ss1)
	checkedExcelData = append(checkedExcelData, ss2)

	fmt.Println(checkedExcelData)
}

func test12() {
	var ss1 []string

	ss1 = append(ss1, "1111")
	ss1 = append(ss1, "22222")
	ss1 = append(ss1, "3333")

	/*
	excelRowNumber := len(ss1)
	ss1 = ss1[4:excelRowNumber]
	fmt.Println(ss1)
	*/


	for i := range ss1{
		fmt.Println(i)
	}

}

func test13()  {
	str := `注：红色字体为必填项，紫色为供应商必填项
	设备代码：机器的IMEI或序列号或机器编码，或可以确认机器信息的本身的其他编码
	物流公司，填写对应的代码编号
	顺丰快递 1； 宅急送 2； EMS 3； EMS落地配 4； 百世快递 5； 圆通快递 6； 韵达快递 7； 中通快递 8； 如风达 9； 
	芝麻开门 10； 万象快递 11； D速快递 12； 山西建华 13； 东骏快递 14； 晟邦快递 15； 黄马甲快递 16； 苏宁快递 17；
	其他 100
	还机地址，填写对应的代码编号
	深圳回收宝（丁德辉）3；北京拿趣用（初漫）4；`

	fmt.Println(str)

	a:=123

	if a == 123 {
		fmt.Println("ok")
	}
}

func oldTest()  {
	//test1()
	//test2()
	//test3()

	//s := test4()
	//fmt.Println(s)

	/*
		test6()
		re:=validate.VerifyMobile("17710750086")
		if  re == false {
			fmt.Println("false")
		}else{
			fmt.Println("true")
		}*/


	//	test7()

	/*
		msg := fmt.Sprintf("%d",int64(123))
		fmt.Println(msg)
	*/

	/*
		num := test8()
		fmt.Println(num)
	*/

	//test9()

	//test11()
}

func main() {

	//oldTest()

	test12()

	/*
	for i:=0; i<100;i++  {
		re := rand.Intn(2)
		fmt.Println(re)
	}
	*/
}
