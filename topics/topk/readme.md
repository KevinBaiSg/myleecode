#TopK

## 问题
从 `arr[1, n]` 这 `n` 个数中，找出最大的 `k` 个数，这就是经典的 TopK 问题。

Example： 
```text
从 arr[1, 12]={5,3,7,1,8,2,9,4,7,2,6,6} 这 n=12 个数中，找出最大的 k=5 个。
```

### 解决方法
#### solution1
思路：暴力排序，返回最大 K 个数

#### solution2
思路：局部排序
优化冒泡排序，因为我们只需要找到最大的 K 个即可，不需要全部排序

#### solution3
思路：堆排序
借助对排序算法，构建最小堆，然后比较堆顶元素，如果大于其堆顶则替换，最后的堆元素即为题解，
由于最小堆可以使用数组实现，也就是说可以在 `nums` 原位置实现整个算法。
所以时间复杂度为最小堆的排序时间复杂度 O(n log n), 空间复杂度为 O(1)

#### solution4
思路：QuickSelect 算法
其实就是修改快排的算法

## LeetCode 类似问题
[215. Kth Largest Element in an Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)
[324. Wiggle Sort II](https://leetcode.com/problems/wiggle-sort-ii/)

## 参考
[1. 拜托，面试别再问我TopK了！！！](https://mp.weixin.qq.com/s/FFsvWXiaZK96PtUg-mmtEw)
