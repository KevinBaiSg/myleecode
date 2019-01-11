/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package addtwonumbers

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var retL *ListNode
	retLC := retL
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		l1V, l2V := 0, 0
		if l1 != nil {
			l1V = l1.Val
		}
		if l2 != nil {
			l2V = l2.Val
		}

		sum := l1V + l2V + carry
		if sum > 9 {
			carry = 1
			sum = sum - 10
		} else {
			carry = 0
		}

		if retL == nil {
			retL = &ListNode{Val: sum, Next: nil}
			retLC = retL
		} else {
			retLC.Next = &ListNode{Val: sum, Next: nil}
			retLC = retLC.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return retL
}
