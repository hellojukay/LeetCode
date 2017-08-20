package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head = new(ListNode)
	var p = head
	for {
		if l1 == nil && l2 == nil {
			p.Next = nil
			break
		} else {
			p.Next = new(ListNode)

		}
		if l1 == nil {
			p = p.Next
			p.Val = l2.Val
			l2 = l2.Next
			p.Next = new(ListNode)
			continue
		}
		if l2 == nil {
			p = p.Next
			p.Val = l1.Val
			l1 = l1.Next
			continue
		}
		if l1.Val < l2.Val {
			p = p.Next
			p.Val = l1.Val
			l1 = l1.Next
		} else {
			p = p.Next
			p.Val = l2.Val
			l2 = l2.Next
		}
	}
	return head.Next
}

func main() {
	var l1 = new(ListNode)
	l1.Val = 1
	l1.Next = nil
	var l2 = new(ListNode)
	l2.Val = 2
	l2.Next = nil
	var l3 = mergeTwoLists(l2, l1)
	for {
		if l3 == nil {
			break
		}
		fmt.Println(l3.Val)
		l3 = l3.Next
	}
}
