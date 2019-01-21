//package Intersection_of_Two_Linked_Lists
package main

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

/*
1.暴力方法：for 嵌套
2.枚举一条链表，记录地址到set中，然后枚举另一条链表，set 中查找是否存在。
3.把一条短的链表尾部接入另一条的头部，问题转换为求闭环的问题。
4.计算两条链表的长度，求逆差后，在同等长度位置开始枚举链表，查找相同值。
*/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	Pa := headA
	Pb := headB
	var Pc *ListNode
	defer func() {
		if Pc != nil {
			Pc.Next = nil
		}
	}()

	for Pa.Next != nil && Pb.Next != nil {
		Pa = Pa.Next
		Pb = Pb.Next
	}

	if Pa.Next == nil {
		Pa.Next = headB
		Pc = Pa
		Pa = headA
	} else {
		Pb.Next = headA
		Pc = Pb
		Pa = headB
	}

	slow := Pa
	fast := Pa.Next

	for slow != fast {
		if slow != nil || fast != nil || fast.Next != nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = Pa
	fast = fast.Next
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main() {
	head := ListNode{Val: 8, Next: &ListNode{4, &ListNode{5, nil}}}
	headA := ListNode{1, &head}
	headB := ListNode{1, &head}
	getIntersectionNode(&headA, &headB)
}
