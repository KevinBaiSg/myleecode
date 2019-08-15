
### 问题描述
23. Merge k Sorted Lists
Hard

Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:
``` 
Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6
``` 

### 解决方法
#### solution1
思路：分治和递归

可以把lists分治来处理，然后合并到一起，这样就转化到合并两个list的问题，本例使用了递归的方式来处理。
