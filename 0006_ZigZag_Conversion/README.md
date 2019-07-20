
### 问题描述
```
6. ZigZag Conversion
Medium
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
```  

### 解决方法
#### solution1
思路：
看到本地本能的思路就是暴力方法，脑海中可以先构建一个二维表，然后枚举 s，
把响应的字符填充到表中，然后从表中横向就可以组装成想要的结果。

很显然这是一个暴力方法，时间复杂度和空间复杂度也会因为二维表增加不少。但是从这个方法中我们可以考虑压缩二维表。
其实构建而且表的目的其实是在构建 `行`。只要我们把字符填到对的行中即可。

所以在解题时，当然先考虑的就是边界条件。
然后根据每个字符的走向（向下还是向上）以及应该在的行，就可以把字符填到对的位置了。

最后可以计算时间复杂度和空间复杂度都为 `O(s)` 
#### solution2
思路：
其实对于一个理工男来说，看到这题的另一个直觉就是应该会有公式的存在。
如果能找到公式，那绝对是最优方式。当然这种方式不推荐在有时间要求的情况下使用。



