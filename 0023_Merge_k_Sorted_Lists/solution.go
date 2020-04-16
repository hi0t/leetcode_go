package solution

import . "leetcode/util"

func mergeKLists(lists []*ListNode) *ListNode {
	res := &ListNode{}

	for _, l := range lists {
		curr := res
		l0 := res.Next
		for l0 != nil && l != nil {
			if l0.Val > l.Val {
				curr.Next = l
				l = l.Next
			} else {
				curr.Next = l0
				l0 = l0.Next
			}
			curr = curr.Next
		}

		for l0 != nil {
			curr.Next = l0
			l0 = l0.Next
			curr = curr.Next
		}

		for l != nil {
			curr.Next = l
			l = l.Next
			curr = curr.Next
		}
	}

	return res.Next
}
