package main

// 0603

func maxRepOpt1(text string) int {
	// 统计每个字符出现的次数
	cnt := [26]int{}
	for _, c := range text {
		cnt[c-'a']++
	}

	n := len(text)
	ans := 0

	// 以每个字符为中心，向左右扩展寻找单字符重复子串
	for i := 0; i < n; {
		j := i
		for j < n && text[j] == text[i] {
			j++
		}

		// 计算以该字符为中心的最长单字符重复子串长度
		l, r := j-i, 0
		k := j + 1
		for k < n && text[k] == text[i] {
			k++
		}
		r = k - j - 1

		// 如果出现次数最多的字符出现次数等于总字符数，特殊处理
		if r > 0 && cnt[text[i]-'a'] == l+r {
			r--
		}

		// 取以不同字符为中心的最长单字符重复子串中的最大值
		ans = max(ans, min(l+r+1, cnt[text[i]-'a']))

		i = j
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
