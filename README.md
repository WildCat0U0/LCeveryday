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


### ---2023.06.05---
#### [LeetCode 2460. 对数组执行操作](https://leetcode.cn/problems/apply-operations-to-an-array/)
##### 题目描述
``简单``
```
给你一个下标从 0 开始的数组 nums ，数组大小为 n ，且由 非负 整数组成。
你需要对数组执行 n - 1 步操作，其中第 i 步操作（从 0 开始计数）要求对 nums 中第 i 个元素执行下述指令：
如果 nums[i] == nums[i + 1] ，则 nums[i] 的值变成原来的 2 倍，nums[i + 1] 的值变成 0 。否则，跳过这步操作。
在执行完 全部 操作后，将所有 0 移动 到数组的 末尾 。
例如，数组 [1,0,2,0,0,1] 将所有 0 移动到末尾后变为 [1,2,1,0,0,0] 。
返回结果数组。
注意 操作应当 依次有序 执行，而不是一次性全部执行。
```

`实现思路`
>对输入数组进行遍历，如果当前数字等于0，那么跳过，num_0++，如果当前数字等于后面的数字，那么当前数字变成2倍，i++，直接跳过下一个数字，num_0++。最后将num_0个0放到数组的末尾即可。
>文档中实现了两种方法，第二种方法简单明了，耗用内存小，两个代码时间复杂度都是O(n)，但是相比之下第二种方式耗时更少，内存更少。

  -第二种方法
  ![img.png](202306/img/img.png) 

  -第一种方法
  ![img_1.png](202306/img/img_1.png)

### ---2023.06.06---
#### [LeetCode 2352.相等行列对](https://leetcode.cn/problems/equal-row-and-column-pairs)
##### 题目描述
``中等``
```
给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
```
``解题思路``
> 利用map和Sprintf将每一行和每一列的数字拼接成字符串，然后将字符串作为key，出现的次数作为value，最后遍历map，将结果一样的内容结果加一

### ---2023.06.07---
#### [LeetCode 2611.老鼠与奶酪](https://leetcode.cn/problems/mice-and-cheese/)
##### 题目描述
``中等``
`动态规划`
```
有两只老鼠和 n 块不同类型的奶酪，每块奶酪都只能被其中一只老鼠吃掉。

下标为 i 处的奶酪被吃掉的得分为：

如果第一只老鼠吃掉，则得分为 reward1[i] 。
如果第二只老鼠吃掉，则得分为 reward2[i] 。
给你一个正整数数组 reward1 ，一个正整数数组 reward2 ，和一个非负整数 k 。

请你返回第一只老鼠恰好吃掉 k 块奶酪的情况下，最大 得分为多少。
```
``解题思路``
>使用贪心算法和排序，使用数组存储两个数组的差值，sum表示第二个数组的和，取出差值最大的k个值，然后把差值加到sum中就是最大值
>也可以使用优先排序的办法，不过这里需要维护的值变成了两个，一个是奶酪的值，一个是数组的长度以及排序，使用时间换空间。但是由于在检索的时候检索的数据从N个变成了k个，所以在时间上也有所优化。
> 第二种算法的代码是
```
type IntHeap []int
 
func (h IntHeap) Less(i, j int) bool {
    return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
    *h = append(*h, x.(int))
}
func (h IntHeap) Len() int {
    return len(h)
}
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n - 1]
    *h = old[:n - 1]
    return x
}

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
    ans := 0
    n := len(reward1)
    pq := &IntHeap{}
    heap.Init(pq)
    for i := 0; i < n; i++ {
        ans += reward2[i]
        diff := reward1[i] - reward2[i]
        heap.Push(pq, diff)
        if pq.Len() > k {
            heap.Pop(pq)
        }
    }
    for pq.Len() > 0 {
        ans += heap.Pop(pq).(int)
    }
    return ans
}

```
### ---2023.06.08---
#### [LeetCode 1240.铺瓷砖](https://leetcode.cn/problems/tiling-a-rectangle-with-the-fewest-squares/)
##### 题目描述
<div style="color: red;">
``困难``
</div>

```
你是一位施工队的工长，根据设计师的要求准备为一套设计风格独特的房子进行室内装修。

房子的客厅大小为 n x m，为保持极简的风格，需要使用尽可能少的 正方形 瓷砖来铺盖地面。

假设正方形瓷砖的规格不限，边长都是整数。

请你帮设计师计算一下，最少需要用到多少块方形瓷砖？
```
``解题思路``
>动态规划问题，可以使用回溯法。


### ---2023.06.09---
#### 偷懒了