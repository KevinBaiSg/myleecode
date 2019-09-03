
### 问题描述
423. Reconstruct Original Digits from English
Medium

Given a non-empty string containing an out-of-order English representation of digits 0-9, output the digits in ascending order.

Note:
Input contains only lowercase English letters.
Input is guaranteed to be valid and can be transformed to its original digits. That means invalid inputs such as "abc" or "zerone" are not permitted.
Input length is less than 50,000.
Example 1:

```text
Input: "owoztneoer"

Output: "012"
```

Example 2:

```text
Input: "fviefuro"

Output: "45"
```

### 解决方法
#### solution1
思路：观察字符与单词的对应关系

   /*
    zero
    one 
    two 
    three
    four
    five
    six
    seven
    eight
    nine
    */
    
//     z -> zero
//     w -> two
//     u -> four
//     x -> six
//     g -> eight

//     h -> three eight 
//     f -> four five
//     s -> six seven
    
//     o -> zero one two four
//     i -> five six eight nine 