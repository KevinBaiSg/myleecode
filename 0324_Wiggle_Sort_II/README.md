
### 问题描述
324. Wiggle Sort II
Medium

Given an unsorted array nums, reorder it such that nums[0] < nums[1] > nums[2] < nums[3]....

Example 1:

```text
Input: nums = [1, 5, 1, 1, 6, 4]
Output: One possible answer is [1, 4, 1, 5, 1, 6].
```

Example 2:

```text
Input: nums = [1, 3, 2, 2, 3, 1]
Output: One possible answer is [2, 3, 1, 3, 1, 2].
```

Note:
You may assume all input has valid answer.

Follow Up:
Can you do it in O(n) time and/or in-place with O(1) extra space?

### 解决方法
#### solution1
思路：QuickSort/HeapSort + three-way-partition
但是排序的时间复杂度已经是 O(n log n) 了，没办法满足题目 O(n) 要求，并且本题也不需要完全排序。

#### solution2
思路：QuickSelect + Three-way-partition

通过 QuickSelect 算法选择出中间数字，然后通过 Three-way-partition 算法转换相应的位置即可。

QuickSelect 的评价时间复杂度为 O(n)，最差为 O(n2).
