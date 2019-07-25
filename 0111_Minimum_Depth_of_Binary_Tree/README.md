
### 问题描述
```
111. Minimum Depth of Binary Tree
Easy

Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.
```  

### 解决方法
思路：还是常规的思路 BFS DFS(stack 或者 递归)
本例使用 DFS 递归方式实现
问题 0104 使用了 BFS 方式实现