package solution

import . "leetcode/util"

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
			continue
		}
		curr = curr.Next
	}

	return head
}
