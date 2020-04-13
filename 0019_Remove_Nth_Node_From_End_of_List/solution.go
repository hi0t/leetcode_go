package solution

import . "leetcode/util"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	l := dummy
	r := dummy

	if n <= 0 {
		return head
	}

	for i := 0; i < n+1; i++ {
		r = r.Next
	}

	for r != nil {
		l = l.Next
		r = r.Next
	}

	l.Next = l.Next.Next

	return dummy.Next
}
