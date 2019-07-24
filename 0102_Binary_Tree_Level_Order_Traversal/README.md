
### 问题描述
```
102. Binary Tree Level Order Traversal
Medium

Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
```  

### 解决方法
#### solution1
思路：BFS
这题应该都会直接想到广度遍历，只是要在遍历过程中解决层的问题。
所以可以在遍历里边加一层，这一层主要负责处理这个 `层`

#### solution1
思路：DFS
DFS 常规的思路是使用递归方式来解题，由于我们还是会面临一个level的问题，所以需要增加一个辅助的函数来传递level