
### 问题描述
179. Largest Number
Medium

Given a list of non negative integers, arrange them such that they form the largest number.

Example 1:

```text
Input: [10,2]
Output: "210"
```

Example 2:

```text
Input: [3,30,34,5,9]
Output: "9534330"
```

Note: The result may be very large, so you need to return a string instead of an integer.

### 解决方法
#### solution1
思路：转换成字符后使用排序算法，比较时使用 a + b > b + a 
