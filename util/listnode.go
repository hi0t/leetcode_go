package util

import "strconv"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr ...int) *ListNode {
	var root, curr *ListNode

	for _, a := range arr {
		if root == nil {
			root = new(ListNode)
			root.Val = a
			curr = root
		} else {
			curr.Next = new(ListNode)
			curr.Next.Val = a
			curr = curr.Next
		}
	}

	return root
}

func (l *ListNode) String() string {
	str := ""
	for l != nil {
		str += strconv.Itoa(l.Val)
		l = l.Next
		if l != nil {
			str += " -> "
		}
	}
	return str
}
