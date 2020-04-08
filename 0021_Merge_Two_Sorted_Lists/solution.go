package solution

import . "leetcode/util"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res

	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			curr.Next = l2
			l2 = l2.Next
		} else {
			curr.Next = l1
			l1 = l1.Next
		}
		curr = curr.Next
	}

	for l1 != nil {
		curr.Next = l1
		l1 = l1.Next
		curr = curr.Next
	}

	for l2 != nil {
		curr.Next = l2
		l2 = l2.Next
		curr = curr.Next
	}

	return res.Next
}
