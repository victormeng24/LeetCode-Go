package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d->", head.Val)
		head = head.Next
	}
}

func BuildList(arr []int) *ListNode {
	dummy := &ListNode{Val: -1}
	tmp := dummy
	for _,v := range arr {
		tmp.Next = &ListNode{Val: v}
		tmp = tmp.Next
	}
	return dummy.Next
}
