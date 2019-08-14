package solution1

import (
	. "github.com/KevinBaiSg/myleecode/common"
)

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n == 0 { return head }

	listHead := ListNode {
		Val: 	0,
		Next: 	head,
	}

	start, end := &listHead, &listHead
	for i := 0; i < n + 1; i++ {
		end = end.Next
	}

	for end != nil {
		start = start.Next
		end = end.Next
	}

	start.Next = start.Next.Next

	return listHead.Next
}
