package makeTree

import (
	"strconv"
	"strings"

	. "github.com/KevinBaiSg/myleecode/common"
	"github.com/KevinBaiSg/myleecode/common/collections"
)

func MakeTree(tree string) *TreeNode {
	tree = strings.TrimSpace(tree)
	tree = strings.Trim(tree, "[")
	tree = strings.Trim(tree, "]")
	nodes := strings.Split(tree, ",")

	queue := collections.Queue{}
	queue.NewQueue()

	i, e := strconv.Atoi(nodes[0])
	if e != nil { return nil }

	root := TreeNode{ i,nil,nil }
	queue.Enqueue(&root)

	index := 1
	for !queue.IsEmpty() {
		node := queue.Dequeue()
		if index < len(nodes){
			node.Left = s2node(nodes[index])
		} else {
			node.Left = nil
		}
		if index + 1 < len(nodes) {
			node.Right = s2node(nodes[index + 1])
		} else {
			node.Right = nil
		}
		if node.Left != nil {
			queue.Enqueue(node.Left)
		}
		if node.Right != nil {
			queue.Enqueue(node.Right)
		}
		index = index + 2
	}

	return &root
}

func s2node(val string) *TreeNode {
	if val == "null" { return nil }
	i, e := strconv.Atoi(val)
	if e != nil { return nil }
	return &TreeNode{i, nil, nil}
}