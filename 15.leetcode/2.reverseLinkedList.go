package main

import "fmt"

//https://leetcode.com/problems/reverse-linked-list/solution/
//https://leetcode-cn.com/problems/reverse-linked-list/

//反转一个单链表。
//示例:
//输入: 1->2->3->4->5->NULL
//输出: 5->4->3->2->1->NULL
//进阶:
//你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

type ListNode struct {
	Val  int
	Next *ListNode
}

//方法1:循环版
func reverseList1(head *ListNode) *ListNode {
	cur:= head
	var prev *ListNode = nil

	for ;cur!=nil;  {
		tmpcur := cur.Next
		cur.Next= prev
		prev = cur
		cur = tmpcur
	}
	return prev
}


//方法1:递归版
func reverseList2(head *ListNode) *ListNode {

	if head ==nil || head.Next==nil{
		return head
	}
	p := reverseList2(head.Next)
	head.Next.Next = head
	head.Next=nil
	return p
}

func travse(node *ListNode) {
	for ;node!=nil; {
		fmt.Print(node.Val)
		node = node.Next
	}
	fmt.Println()
}

//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL
func main() {
	var head *ListNode
	n5 := ListNode{5, nil}
	n4 := ListNode{4, &n5}
	n3 := ListNode{3, &n4}
	n2 := ListNode{2, &n3}
	n1 := ListNode{1, &n2}

	travse(&n1)
	head = reverseList1(&n1)
	travse(head)


	n5 = ListNode{5, nil}
	n4 = ListNode{4, &n5}
	n3 = ListNode{3, &n4}
	n2 = ListNode{2, &n3}
	n1 = ListNode{1, &n2}

	travse(&n1)
	head = reverseList2(&n1)
	travse(head)


}
