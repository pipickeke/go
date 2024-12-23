package main

/*
题目：2980. 检查按位或是否存在尾随零

给你一个 正整数 数组 nums 。

你需要检查是否可以从数组中选出 两个或更多 元素，
满足这些元素的按位或运算（ OR）结果的二进制表示中 至少 存在一个尾随零。

例如，数字 5 的二进制表示是 "101"，不存在尾随零，而数字 4 的二进制表示是 "100"，存在两个尾随零。

如果可以选择两个或更多元素，其按位或运算结果存在尾随零，返回 true；否则，返回 false 。
*/

func hasTrailingZeros(num []int) bool {
	even := len(num)
	for _, x := range num {
		even -= x % 2
	}
	return even >= 2
}
