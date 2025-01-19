package main

/**
题目：746. 使用最小花费爬楼梯
标签：动态规划

给你一个整数数组 cost ，其中 cost[i] 是从楼梯第 i 个台阶向上爬需要支付的费用。
一旦你支付此费用，即可选择向上爬一个或者两个台阶。

你可以选择从下标为 0 或下标为 1 的台阶开始爬楼梯。

请你计算并返回达到楼梯顶部的最低花费。

*/

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}
		return min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
	}
	return dfs(n)
}

func minCostClimbingStairs2(cost []int) int {
	n := len(cost)
	memo := make([]int, n+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}
		p := &memo[i]
		if *p != -1 {
			return *p
		}
		res := min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
		*p = res
		return res
	}
	return dfs(n)
}

func minCostClimbingStairs3(cost []int) int {
	n := len(cost)
	ans := make([]int, n+1)
	for i := 2; i <= n; i++ {
		ans[i] = min(ans[i-1]+cost[i-1], ans[i-2]+cost[i-2])
	}
	return ans[n]
}
