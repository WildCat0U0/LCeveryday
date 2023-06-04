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

### ---2023.06.03---
#### [LeetCode 1156. 单字符重复子串的最大长度](https://leetcode.cn/problems/swap-for-longest-repeated-character-substring/)
##### 题目描述
``中等``
```
如果字符串中的所有字符都相同，那么这个字符串是单字符重复的字符串。
给你一个字符串 text，你只能交换其中两个字符一次或者什么都不做，然后得到一些单字符重复的子串。返回其中最长的子串的长度。
```
例如 ：
输入：text = "ababa"
输出：3

**双指针**
一方面是理解为：以单个字母为中心，向两边拓展同一个字符的最长串，直到遇到了不同的字符停止，比如abaaba，以`b`为中心，向左边拓展指针最长长度
为1，向右边最长长度为2，但是必须先查看当前a的个数有没有等于整个字符串的总数，如果没有等于，那么说明中间的b可以找另外的a来进行替换。那么结果就是`l+r+1`
循环以往，直到找到最长的串。

### ---2023.06.04---
#### [LeetCode 2465. 单字符重复子串的最大长度](https://leetcode.cn/problems/number-of-distinct-averages)
##### 题目描述
``简单``
```给你一个下标从 0 开始长度为 偶数 的整数数组 nums 。

只要 nums 不是 空数组，你就重复执行以下步骤：

找到 nums 中的最小值，并删除它。
找到 nums 中的最大值，并删除它。
计算删除两数的平均值。
两数 a 和 b 的 平均值 为 (a + b) / 2 。

比方说，2 和 3 的平均值是 (2 + 3) / 2 = 2.5 。
返回上述过程能得到的 不同 平均值的数目。

注意 ，如果最小值或者最大值有重复元素，可以删除任意一个。
```

`实现思路`
>排序，每次删除头尾，保存平均值，将结果保存到一个map[float]struct{}里面，map是直接去重的，最后返回map的长度

