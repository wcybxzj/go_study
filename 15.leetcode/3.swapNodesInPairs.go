package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func travse(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

func swapPairs(head *ListNode) *ListNode {
	var prev *ListNode
	var tmp *ListNode

	self := ListNode{0, nil}
	prev = &self
	prev.Next = head

	for prev.Next != nil && prev.Next.Next != nil {
		a := prev.Next
		b := a.Next
		prev.Next = b

		tmp = b.Next
		b.Next = a
		a.Next = tmp
		prev = a
	}

	return self.Next
}

//https://leetcode.com/problems/swap-nodes-in-pairs
//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
//示例:
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//给定 1->2->3->4->5, 你应该返回 2->1->4->3->5.
func main() {
	var head *ListNode
	n5 := ListNode{5, nil}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}

	travse(&n1)
	head = swapPairs(&n1)
	travse(head)

}
