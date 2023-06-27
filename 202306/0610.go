package abc

import "sort"

// 思路简单，但是时间复杂度较高，空间复杂度较高
func numSmallerByFrequency(queries []string, words []string) []int {
	var minChar int
	var minCharCnt int
	var queriesCnt []int
	var wordsCnt []int
	for i := range queries {
		minCharCnt = 0
		minChar = 0
		for j := range queries[i] {
			if queries[i][j] < queries[i][minChar] {
				minChar = j
				minCharCnt = 1
			} else if queries[i][j] == queries[i][minChar] {
				minCharCnt++
			}
		}
		queriesCnt = append(queriesCnt, minCharCnt)
	}
	for i := range words {
		minCharCnt = 0
		minChar = 0
		for j := range words[i] {
			if words[i][j] < words[i][minChar] {
				minChar = j
				minCharCnt = 1
			} else if words[i][j] == words[i][minChar] {
				minCharCnt++
			}
		}
		wordsCnt = append(wordsCnt, minCharCnt)
	}
	var res []int
	for i := range queriesCnt {
		cnt := 0
		for j := range wordsCnt {
			if queriesCnt[i] < wordsCnt[j] {
				cnt++
			}
		}
		res = append(res, cnt)
	}
	return res
}

// ----------------------------------------------------------------------------------------------
// 这种方法在空间复杂度上比较优秀，但是时间复杂度仍然很高
func f(s string) int {
	cnt := 0
	ch := 'z'
	for _, c := range s {
		if c < ch {
			ch = c
			cnt = 1
		} else if c == ch {
			cnt++
		}
	}
	return cnt
}

func numSmallerByFrequencyLeetCodeOfficial(queries []string, words []string) []int {
	count := make([]int, 12)
	for _, s := range words {
		count[f(s)] += 1
	}
	for i := 9; i >= 1; i-- {
		count[i] += count[i+1]
	}
	res := make([]int, len(queries))
	for i, s := range queries {
		res[i] = count[f(s)+1]
	}
	return res
}

// ----------------------------------------------------------------------------------------------
// 这种方法在空间复杂度上比较高，但是时间复杂度很低

func numSmallerByFrequencyLeetCodeUsers(queries []string, words []string) (ans []int) {
	f := func(s string) int {
		cnt := [26]int{}
		for _, c := range s {
			cnt[c-'a']++
		}
		for _, x := range cnt {
			if x > 0 {
				return x
			}
		}
		return 0
	}
	n := len(words)
	nums := make([]int, n)
	for i, w := range words {
		nums[i] = f(w)
	}
	sort.Ints(nums)
	for _, q := range queries {
		x := f(q)
		ans = append(ans, n-sort.SearchInts(nums, x+1))
	}
	return ans
}
