package main

/**
题目：898. 子数组按位或操作
标签：logtrick

给定一个整数数组 arr，返回所有 arr 的非空子数组的不同按位或的数量。

子数组的按位或是子数组中每个整数的按位或。含有一个整数的子数组的按位或就是该整数。

子数组 是数组内连续的非空元素序列。

*/

func subarrayBitwiseORs(arr []int) int {
	set := map[int]bool{}
	for i, x := range arr {
		set[x] = true
		for j := i - 1; j >= 0 && (arr[j]|x) != arr[j]; j-- {
			arr[j] |= x
			set[arr[j]] = true
		}
	}
	return len(set)
}
