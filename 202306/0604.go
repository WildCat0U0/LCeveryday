package main

import "sort"

func distinctAverages(nums []int) int {
	res := make(map[float32]struct{}, 0)
	sort.Ints(nums) // 排序
	var result float32
	for i := 0; i < len(nums)/2; i++ {
		result = (float32(nums[i] + nums[len(nums)-1-i])) / 2
		res[result] = struct{}{}
	}
	return len(res)
}
