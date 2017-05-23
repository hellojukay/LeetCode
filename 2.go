package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p := new(ListNode)
	q := p
	s := 0
	for {
		if l1 != nil {
			s = s + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			s = s + l2.Val
			l2 = l2.Next
		}
		p.Val = s % 10
		s = s / 10
		if (l1 == nil) && (l2 == nil) {
			if s != 0 {
				p.Next = new(ListNode)
				p = p.Next
				p.Val = s
			}
			p.Next = nil
			break
		}
		p.Next = new(ListNode)
		p = p.Next
	}
	return q
}

func toLink([]int) *ListNode {

}

func toArray(l *ListNode) []int {

}
