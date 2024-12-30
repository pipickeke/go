package main

/*
题目：3133. 数组最后一个元素的最小值

给你两个整数 n 和 x 。你需要构造一个长度为 n 的 正整数
数组 nums ，对于所有 0 <= i < n - 1 ，满足 nums[i + 1] 大于 nums[i] ，
并且数组 nums 中所有元素的按位 AND 运算结果为 x 。

返回 nums[n - 1] 可能的 最小 值。

*/

func minEnd(n, x int) int64 {
	n--
	i, j := 0, 0
	for n>>j > 0 {
		if x>>i&1 == 0 {
			x |= (n >> j & 1) << i
			j++
		}
		i++
	}
	return int64(x)
}
