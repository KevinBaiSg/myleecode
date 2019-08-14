
### 问题描述

Letter Combinations of a Phone Number      
Medium  

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.   

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters. 

![](http://upload.wikimedia.org/wikipedia/commons/thumb/7/73/Telephone-keypad2.svg/200px-Telephone-keypad2.svg.png)

Example:    

Input: "23" 
``` 
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:   
```     

Although the above answer is in lexicographical order, your answer could be in any order you want.  
  

### 解决方法    
#### solution1  
思路：回溯

开始解这题的时候觉得挺简单，但是在写代码的时候犯糊涂了。总结起来就是本题解答方法用人类思维方式很好回答，
但是用计算机实现的时候还是需要严格按照计算机的算法数据结构来实现。   

本题其实就是树的深度搜索实现，具体介绍可以看代码中的注释。
