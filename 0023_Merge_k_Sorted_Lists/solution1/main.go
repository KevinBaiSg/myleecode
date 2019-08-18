package solution1

import (
	. "github.com/KevinBaiSg/MyLeetCode/common"
)

func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil { return nil }

	length := len(lists)

	if length == 0 { return nil }
	if length == 1 { return lists[0] }

	half := length / 2

	left := mergeKLists(lists[:half])
	right := mergeKLists(lists[half:])
	return mergeTwoLists(left, right)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil { return list2 }
	if list2 == nil { return list1 }

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list1, list2.Next)
	return list2
}
