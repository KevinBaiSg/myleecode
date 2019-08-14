
### 问题描述
19. Remove Nth Node From End of List
Medium

Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

### 解决方法
#### solution1
思路：标尺法（我命名的。。。）

这道题唯一的难点在于 `one pass`，所以我们可以使用双指针展开一个尺子，量出 `n` 长。在写代码的时候我手动增加了一个 list head，
主要是为了后序要处理边界的问题，这也是一种常用的链表处理方式，可以方便不少。
