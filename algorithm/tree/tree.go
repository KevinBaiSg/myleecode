package tree

import (
	. "github.com/KevinBaiSg/myleecode/common"
	. "github.com/KevinBaiSg/myleecode/common/collections"
)

func preOrderTraversalRecursion(root *TreeNode, result *[]int) {
	if root == nil || result == nil { return }

	*result = append(*result, root.Val) // 相当于访问节点
	preOrderTraversalRecursion(root.Left, result)
	preOrderTraversalRecursion(root.Right, result)
}

func inOrderTraversalRecursion(root *TreeNode, result *[]int) {
	if root == nil || result == nil { return }

	inOrderTraversalRecursion(root.Left, result)
	*result = append(*result, root.Val) // 相当于访问节点
	inOrderTraversalRecursion(root.Right, result)
}

func postOrderTraversalRecursion(root *TreeNode, result *[]int) {
	if root == nil || result == nil { return }

	postOrderTraversalRecursion(root.Left, result)
	postOrderTraversalRecursion(root.Right, result)
	*result = append(*result, root.Val) // 相当于访问节点
}

func preOrderTraversal(root *TreeNode) (result []int) {
	if root == nil { return []int{} }

	stack := Stack{}
	stack.NewStack()

	pNode := root
	for !stack.IsEmpty() || pNode != nil {
		if pNode != nil {
			result = append(result, pNode.Val) // 访问节点
			stack.Push(pNode)
			pNode = pNode.Left
		} else {
			node := stack.Pop()
			pNode = node.Right
		}
	}

	return
}

func inOrderTraversal(root *TreeNode) (result []int) {
	if root == nil { return []int{} }

	stack := Stack{}
	stack.NewStack()

	pNode := root
	for pNode != nil || !stack.IsEmpty() {
		if pNode != nil {
			stack.Push(pNode)
			pNode = pNode.Left
		} else {
			node := stack.Pop()
			result = append(result, node.Val)
			pNode = node.Right
		}
	}

	return
}

func postOrderTraversal(root *TreeNode) (result []int) {
	if root == nil { return []int{} }

	/*
	1.增加标志位标识右节点是否访问过。
	2.方法二：转化问题，左->右->根 = 翻转(根->右->左)
	*/
	stack := Stack{}
	stack.NewStack()

	pNode := root
	/*
		pLastVisited 用于标记上次访问是否为该右节点
	*/
	var pLastVisited *TreeNode  = nil

	for pNode != nil || !stack.IsEmpty() {
		if pNode != nil {
			stack.Push(pNode)
			pNode = pNode.Left
		} else {
			node := stack.Peek()
			if node.Right == nil || node.Right == pLastVisited {
				// node.Right == nil 没有右节点，继续出栈
				// node.Right == pLastVisited 代表右节点刚刚访问过，继续出栈
				result = append(result, node.Val)
				stack.Pop()
				pLastVisited = node
			} else {
				pNode = node.Right // 可以看做子树
			}
		}
	}
	return
}


func layerTraver(root *TreeNode) (result []int) {
	if root == nil { return []int{} }

	queue := Queue{}
	queue.NewQueue()

	queue.Enqueue(root)

	for !queue.IsEmpty() {
		n := queue.Size()
		for i := 0; i < n; i++ {
			node := queue.Dequeue()
			result = append(result, node.Val) //访问节点

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}
	return
}