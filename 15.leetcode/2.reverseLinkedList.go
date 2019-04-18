package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reverseList(head *ListNode) *ListNode {
//
//}

func create(arr1 []int) *ListNode {
	var prev *ListNode
	var head ListNode
	for index, value := range arr1 {
		//fmt.Println(index, value)
		if index == 0 {
			head := ListNode{Val: value, Next: nil}
			prev = &head
		} else {
			node := &ListNode{Val: value, Next: nil}
			prev.Next = node
			prev = node
		}
	}
	return &head
}

func travse(node *ListNode) {
	for (*node).Next != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

//没写完去学nginx了
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	head := create(arr1)
	travse(head)
}
