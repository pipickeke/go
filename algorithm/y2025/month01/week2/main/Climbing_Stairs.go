package main

/**
题目：70. 爬楼梯
标签：动态规划

假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
*/

// 超时1
func climbStairs(n int) int {
	if n < 2 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)

}

func climbStairs2(n int) int {
	memo := make([]int, n+1)
	var dfs func(n int) int
	dfs = func(n int) int {
		if n < 2 {
			return 1
		}
		p := &memo[n]
		if *p == 0 {
			*p = dfs(n-1) + dfs(n-2)
		}
		return *p
	}
	return dfs(n)
}
