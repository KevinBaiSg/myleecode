
### 问题描述
```
1038. Binary Search Tree to Greater Sum Tree
Medium

Given the root of a binary search tree with distinct values, modify it so that every node has a new value equal to the sum of the values of the original tree that are greater than or equal to node.val.

As a reminder, a binary search tree is a tree that satisfies these constraints:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.
 

Example 1:



Input: [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
 

Note:

The number of nodes in the tree is between 1 and 100.
Each node will have value between 0 and 100.
The given tree is a binary search tree.
```  

### 解决方法
#### solution1
思路：
这道题我中文看了很多遍，英文看了很多遍最后还是没有搞懂什么意思，
也看了别人解题的方法还是没有搞懂，最后看到一个和我有类似经历的人的描述最后算是搞懂题了。

这是那人的原话
``` 
我发现我输就输在阅读理解上。 题意解析： 首先我第一遍几乎没看懂到底这个题想干嘛？ 
然后进一步：为什么修改后的二叉树每个子树右节点比左边节点的值要小？（人家也没讲修改后必须也要是BST啊，看来是我智障） 
接下来是这句话“使每个节点 node 的新值等于原树中大于或等于 node.val 的值之和” （我脑子里是一团☁️。） 
看来几遍终于懂了：修改当前节点的值a，使得新值x 等于 原树中 大于或者等于 a 的值 之和。 
所以右中左遍历，每到一个节点，将原来遍历过的值的和替换掉原来的值就OK了。 
（感觉这道题所谓的中等难度 很大一部分是在阅读理解的难度上）

作者：hardy-5
链接：https://leetcode-cn.com/problems/two-sum/solution/dui-ti-yi-de-li-jie-by-hardy-5/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
``` 
其实这题可以把 tree 转化成 array  [4,1,6,0,2,5,7,3,8], 在计算 `7` 位置的新值x时，其实就是 x = [sum(>=7的所有值)]

如果能理解到这，这题的思路就出来了，就是遍历 tree，新值就是比他大的值得和，至于怎么遍历就是需要些技巧了。

考虑到 bst 的特质，就会考虑使用中序遍历，但是要先从右子树开始，因为那边是最大。这样基本思路就形成了：中序遍历 + 右子树开始 + 保存一个和

由于 Go 语言的问题，我们需要借助一个函数来传递 `和`。