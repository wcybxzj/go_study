package queue

//Golang: interface{}
//像c++中的template
//像c中的void*

//限定方法1:在声明时候限定Queue中的元素类型
//type Queue []int
type Queue []interface{}

//API1:
//因为函数中接受者为指针
//所以无论调用时候接受者是值还是指针都会强转成指针
func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

//API2:
//限定方法2:在借口上限制只能是int
func (q *Queue) Push2(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop2() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) //Golang垃圾语法
}

//API3:
//限定方法3:在函数内部对类型进行限定
func (q *Queue) Push3(v interface{}) {
	*q = append(*q, v.(int))
}
