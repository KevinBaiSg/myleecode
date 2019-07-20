package reverselinkedlist

/**
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseListrecursively(head *ListNode) *ListNode {

// 	// terminator: if nil return
// 	if head.Next == nil {
// 		return head
// 	}
// 	// process reversel sub list

// 	// drill down next sub list

// 	// clear status no

// }

func reverseListiteratively(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}
