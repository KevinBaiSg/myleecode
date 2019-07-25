
### 问题描述
```
104. Maximum Depth of Binary Tree
Easy

Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
```  

### 解决方法
思路：还是常规的思路 BFS DFS(stack 或者 递归)
本例使用 BFS 实现
问题 0111 使用了 DFS 递归方式实现