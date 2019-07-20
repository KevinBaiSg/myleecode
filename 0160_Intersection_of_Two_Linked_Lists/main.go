package Intersection_of_Two_Linked_Lists

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
3.两条链表同时移动，当一条到尾部的时候，该节点next接到短链的头，长链继续移动，
	原理是，同时移动回到原点可以保证两条链表同步了
4.计算两条链表的长度，求逆差后，在同等长度位置开始枚举链表，查找相同值。
*/
func getIntersectionNode2(headA, headB *ListNode) *ListNode {

	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
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

	if Pa.Next == nil && Pb.Next == nil {
		Pa = headA
		Pb = headB
	} else if Pa.Next == nil {
		Pc = Pa
		Pa = headB
		Pb = Pb.Next
		Pc.Next = headA
	} else {
		Pc = Pb
		Pb = headA
		Pa = Pa.Next
		Pc.Next = headB
	}

	for Pa != nil && Pb != nil {
		if Pa == Pb {
			return Pa
		}
		Pa = Pa.Next
		Pb = Pb.Next
	}
	return nil
}

//func main() {
//	head := ListNode{2, &ListNode{4, nil}}
//	headA := ListNode{0, &ListNode{9, &ListNode{1, &head}}}
//	headB := ListNode{3, &head}
//	getIntersectionNode3(&headA, &headB)
//}
