package queue

type Queue []int

//因为函数中接受者为指针
//所以无论调用时候接受者是值还是指针都会强转成指针
func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
