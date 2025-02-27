package main

/**
题目：1542. 找出最长的超赞子字符串

给你一个字符串 s 。请返回 s 中最长的 超赞子字符串 的长度。

「超赞子字符串」需满足满足下述两个条件：

该字符串是 s 的一个非空子字符串
进行任意次数的字符交换后，该字符串可以变成一个回文字符串

*/

func longestAwesome(s string) int {
	const NUMS = 10
	n := len(s)
	pos := [1 << NUMS]int{}

	for i := range pos {
		pos[i] = n
	}
	pos[0] = -1

	pre := 0
	ans := 0
	for i, c := range s {
		pre ^= 1 << (c - '0')
		for d := 0; d < NUMS; d++ {
			ans = max(ans, i-pos[pre^(1<<d)])
		}
		ans = max(ans, i-pos[pre])
		if pos[pre] == n {
			pos[pre] = i
		}
	}
	return ans
}
