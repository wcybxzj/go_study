package main

import "fmt"

//测试1:
//知识点1:append数据到切片,系统会重新分配更大的底层数据给切片做映射
//知识点2:append后一定要用变量来接受,因为当append据到原切片,
//原切片的ptr和cap都会变化

//分片是数据的视图可如果修改了分片就会引发COW会单独为分片分配空间,
//而原数组不变

//输出:
//s3,s4,s5= [5 6 10] [5 6 10 11] [5 6 10 11 12]
//arr= [0 1 2 3 4 5 6 10]
func test1() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] //2,3,4,5
	s2 := s1[3:5]  //5,6

	//append后s3,s4,s5不再是arr的视图,
	//COW时候会生成一个新的匿名arr,让s3,s4,s5来做视图
	s3 := append(s2, 10) //?
	s4 := append(s3, 11) //?
	s5 := append(s4, 12) //?
	fmt.Println("s3,s4,s5=", s3, s4, s5)
	fmt.Println("arr=", arr)
}

func main() {
	test1()
}
