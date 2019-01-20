package Remove_Linked_List_Elements

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}

	return dummyHead.Next
}

func removeElements2(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	result := removeElements2(head.Next, val)
	if head.Val == val {
		return result
	} else {
		head.Next = result
		return head
	}
}
