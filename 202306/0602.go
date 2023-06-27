package abc

import "sort"

// Path: 202306\0602.go
// 中文
// 解法1：前缀和 利用数组表示当前位置之前的元音字母个数 直接输出的就是二维数组两个结果的差值
// 时间复杂度 O(n+m)其中n为单词的数量，m为单词的最大长度

func vowelStrings0(words []string, queries [][]int) []int {
	wordLen := len(words)
	queriesLen := len(queries)

	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	// 二维数组 保存每个单词的前缀和
	perfix := make([]int, wordLen+1)
	for i, w := range words {
		x := 0
		if vowels[w[0]] && vowels[w[len(w)-1]] {
			x = 1
		}
		perfix[i+1] = perfix[i] + x
	}
	ans := make([]int, queriesLen) // 保存结果
	for i, q := range queries {    // 遍历查询数组
		ans[i] = perfix[q[1]+1] - perfix[q[0]] // 直接输出的就是二维数组两个结果的差值
	}
	return ans
}

// 解法2:预处理+二分查找
// 时间复杂O(nqlog n)其中n为单词的数量，q为查询的数量
func vowelStrings1(words []string, queries [][]int) []int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	nums := make([]int, 0)
	for i, w := range words {
		if vowels[w[0]] && vowels[w[len(w)-1]] {
			nums = append(nums, i)
		}
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		l, r := q[0], q[1]
		ans[i] = sort.SearchInts(nums, r+1) - sort.SearchInts(nums, l)
	}
	return ans
}
