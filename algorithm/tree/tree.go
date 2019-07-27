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
	// TODO:
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