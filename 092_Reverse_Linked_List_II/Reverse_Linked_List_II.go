package _92_Reverse_Linked_List_II

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var (
		pre *ListNode
		cur = head
		prestart *ListNode
		end *ListNode
	)

	// Note: 1 ≤ m ≤ n ≤ length of list.
	for i := 0; i < n; i++ {
		if i < m - 1 {
			pre, prestart = cur, cur
			cur = cur.Next
			end = cur
		} else {
			cur.Next, pre, cur = pre, cur, cur.Next
		}
	}

	if m > 1 {
		prestart.Next = pre
		end.Next = cur
		return head
	} else {
		head.Next = cur
		return pre
	}
}