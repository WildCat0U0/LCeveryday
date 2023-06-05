package main

import "fmt"

// 解法1 从左到右遍历，遇到相同的数字就合并，遇到0就记录0的个数，最后补0 速度快，但是内存占用高
func applyOperations(nums []int) []int {
	// nums 输入的数组
	var numberOfZero int = 0
	res := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			numberOfZero++
		} else if nums[i] == nums[i+1] {
			res = append(res, nums[i]*2)
			numberOfZero++
			i++
		} else {
			res = append(res, nums[i])
		}
	}
	if nums[len(nums)-1] != 0 && nums[len(nums)-1] != nums[len(nums)-2] {
		res = append(res, nums[len(nums)-1])
	} else if nums[len(nums)-1] == 0 {
		numberOfZero++
	}
	for i := 0; i < numberOfZero; i++ {
		res = append(res, 0)
	}
	return res
}

// 官方解法， 简单明了 在数组本身进行操作，不需要额外的空间，每次遇到0，最后一位始终不动，除非
func applyOperations_1(nums []int) []int {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if i+1 < n && nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}

func main() {
	res := applyOperations_1([]int{1, 2, 2, 2, 1, 1, 1, 1, 1, 3, 3, 2, 1, 3, 4, 5, 0, 0, 0, 1, 2, 3})
	fmt.Println(res)
}
