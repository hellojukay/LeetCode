package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var i = n
	var p = head
	var q = head
	for {
		if i == 0 {
			break
		}
		p = p.Next
		i--
	}
	if p == nil {
		head = head.Next
		return head
	}
	for {
		if p.Next == nil {
			if q.Next.Next == nil {
				q.Next = nil
				break

			}
			q.Next = q.Next.Next
			break

		}
		p = p.Next
		q = q.Next

	}
	return head
}
