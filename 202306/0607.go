package abc

import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n := len(reward1)
	differences := make([]int, n) //map[key]value
	res := 0
	for i := 0; i < n; i++ {
		res += reward2[i]
		differences[i] = reward1[i] - reward2[i]
	}
	sort.Ints(differences)
	for i := 1; i <= k; i++ {
		res += differences[n-i]
	}
	return res
}
