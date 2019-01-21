package Palindrome_Linked_List

//package main

/**
ref: https://leetcode.com/problems/palindrome-linked-list/
*/

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

func reverse(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}
	return pre
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head
	var rHalf *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil { // 偶
		rHalf = slow
	} else { // 奇
		rHalf = slow.Next
	}

	reverseList := reverse(rHalf)
	compareA := head
	compareB := reverseList

	for compareB != nil {
		if compareA.Val != compareB.Val {
			return false
		}
		compareA = compareA.Next
		compareB = compareB.Next
	}

	return true
}

//func main() {
//	isPalindrome(&ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}})
//}
