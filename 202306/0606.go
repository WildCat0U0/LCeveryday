package abc

import "fmt"

func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := make(map[string]int)
	for _, row := range grid {
		cnt[fmt.Sprint(row)]++ // 以字符串的形式存储数组的每一行
	}
	res := 0
	for j := 0; j < n; j++ {
		var arr []int
		for i := 0; i < n; i++ {
			arr = append(arr, grid[i][j]) //把每一列的数字放到数组中
		}
		if val, ok := cnt[fmt.Sprint(arr)]; ok { //如果存在这一行，那么结果就加一
			res += val
		}
	}

	return res
}

//时间复杂度为O(n^2) 空间复杂度为O(n^2)
