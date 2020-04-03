package main

import (
	"fmt"
	. "leetcode/util"
)

func main() {
	l1 := NewListNode(2, 4, 3)
	l2 := NewListNode(5, 6, 4)

	res := addTwoNumbers(l1, l2)

	fmt.Println(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var res, curr *ListNode
	carry := 0

	for l1 != nil || l2 != nil {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		sum += carry

		if sum > 9 {
			sum -= 10
			carry = 1
		} else {
			carry = 0
		}

		if res == nil {
			res = new(ListNode)
			res.Val = sum
			curr = res
		} else {
			curr.Next = new(ListNode)
			curr.Next.Val = sum
			curr = curr.Next
		}
	}

	if carry > 0 {
		curr.Next = new(ListNode)
		curr.Next.Val = carry
	}

	return res
}
