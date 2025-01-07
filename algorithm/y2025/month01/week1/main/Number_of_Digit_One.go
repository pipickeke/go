package main

import "strconv"

/**
题目：233. 数字 1 的个数

给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。


*/

func countDigitOne(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dfs func(int, int, bool) int
	dfs = func(i int, cnt int, islimit bool) (res int) {
		if i == m {
			return cnt
		}

		if !islimit {
			p := &memo[i][cnt]
			if *p >= 0 {
				return *p
			}
			defer func() { *p = res }()
		}
		up := 9
		if islimit {
			up = int(s[i] - '0')
		}

		for d := 0; d <= up; d++ {
			c := cnt
			if d == 1 {
				c++
			}
			res += dfs(i+1, c, islimit && d == up)
		}
		return
	}
	return dfs(0, 0, true)
}
