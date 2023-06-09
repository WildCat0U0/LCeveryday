package main

// 暴力求解
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tilingRectangle(n int, m int) int {
	ans := max(n, m)
	rect := make([][]bool, n)
	for i := 0; i < n; i++ {
		rect[i] = make([]bool, m)
	}

	isAvailable := func(x, y, k int) bool {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				if rect[x+i][y+j] {
					return false
				}
			}
		}
		return true
	}

	fillUp := func(x, y, k int, val bool) {
		for i := 0; i < k; i++ {
			for j := 0; j < k; j++ {
				rect[x+i][y+j] = val
			}
		}
	}

	var dfs func(int, int, int)
	dfs = func(x, y, cnt int) {
		if cnt >= ans {
			return
		}
		if x >= n {
			ans = cnt
			return
		}
		// 检测下一行
		if y >= m {
			dfs(x+1, 0, cnt)
			return
		}
		// 如当前已经被覆盖，则直接尝试下一个位置
		if rect[x][y] {
			dfs(x, y+1, cnt)
			return
		}
		for k := min(n-x, m-y); k >= 1 && isAvailable(x, y, k); k-- {
			// 将长度为 k 的正方形区域标记覆盖
			fillUp(x, y, k, true)
			// 跳过 k 个位置开始检测
			dfs(x, y+k, cnt+1)
			fillUp(x, y, k, false)
		}
	}

	dfs(0, 0, 0)
	return ans
}

// 优化
func tilingRectangle1(n int, m int) int {
	ans := n * m
	filled := make([]int, n)
	var dfs func(i, j, t int)
	dfs = func(i, j, t int) {
		if j == m {
			i++
			j = 0
		}
		if i == n {
			ans = t
			return
		}
		if filled[i]>>j&1 == 1 {
			dfs(i, j+1, t)
		} else if t+1 < ans {
			var r, c int
			for k := i; k < n; k++ {
				if filled[k]>>j&1 == 1 {
					break
				}
				r++
			}
			for k := j; k < m; k++ {
				if filled[i]>>k&1 == 1 {
					break
				}
				c++
			}
			mx := min(r, c)
			for w := 1; w <= mx; w++ {
				for k := 0; k < w; k++ {
					filled[i+w-1] |= 1 << (j + k)
					filled[i+k] |= 1 << (j + w - 1)
				}
				dfs(i, j+w, t+1)
			}
			for x := i; x < i+mx; x++ {
				for y := j; y < j+mx; y++ {
					filled[x] ^= 1 << y
				}
			}
		}
	}
	dfs(0, 0, 0)
	return ans
}
