package solution

import . "leetcode/util"

func reverseKGroup(head *ListNode, k int) *ListNode {
	curr := head
	var tail, h *ListNode

	for curr != nil {
		cnt := 0

		curr = head

		for ; cnt < k && curr != nil; cnt++ {
			curr = curr.Next
		}

		if cnt == k {
			rev := reverseList(head, k)

			if h == nil {
				h = rev
			}

			if tail != nil {
				tail.Next = rev
			}

			tail = head
			head = curr
		}
	}

	if tail != nil {
		tail.Next = head
	}

	if h == nil {
		return head
	}

	return h
}

func reverseList(head *ListNode, k int) *ListNode {
	var h *ListNode
	curr := head
	for ; k > 0; k-- {
		next := curr.Next
		curr.Next = h
		h = curr
		curr = next
	}
	return h
}
