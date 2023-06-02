# LCeveryday
## 2023-06
_____________________________________________________________________________________________
### ---2023.06.02---
#### [LeetCode 2559 统计范围内的原因字符串数](https://leetcode.cn/problems/count-vowel-strings-in-ranges/)
##### 题目描述
``中等``

```
给你一个下标从 0 开始的字符串数组 words 以及一个二维整数数组 queries 。
每个查询 queries[i] = [li, ri] 会要求我们统计在 words 中下标在 li 到 ri 范围内（包含 这两个值）并且以元音开头和结尾的字符串的数目。
返回一个整数数组，其中数组的第 i 个元素对应第 i 个查询的答案。
注意：元音字母是 'a'、'e'、'i'、'o' 和 'u' 。
```
* 这个问题是一个字符串匹配问题，需要找到words数组中所有以元音字母开头和结尾的字符串。为了减少时间复杂度，我们使用前缀和的方法来预处理words
  数组中所有以元音字母开头和结尾的字符串数量，然后对于每个查询，用前缀和的差值来计算该范围内符合条件的字符串数量。
* -具体地，我们首先遍历words数组，统计出每个位置之前以元音字母开头和结尾的字符串数量，这样可以在O(1)时间内计算出任意区间的以元音字母开头和
结尾的字符串数量。然后对于每个查询，我们只需要用前缀和的差值来计算该范围内符合条件的字符串数量。
* 具体地，我们用``prefix[i][j]``表示``words[0][:i]``中以元音字母j为结尾的字符串数量，然后对于每个查询``[l, r]``，我们计算 temp``[j]``表示words``[l:r+1]``
中以元音字母j为开头的字符串数量，然后使用双重循环计算符合条件的字符串数量。时间复杂度为``O(n+q)``，其中``n``为words数组的长度，``q``为queries数组
的长度。


