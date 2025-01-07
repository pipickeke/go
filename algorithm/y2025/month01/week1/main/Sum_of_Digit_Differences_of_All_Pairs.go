package main

import "strconv"

/**
题目：3153. 所有数对中数位差之和

你有一个数组 nums ，它只包含 正 整数，所有正整数的数位长度都 相同 。

两个整数的 数位差 指的是两个整数 相同 位置上不同数字的数目。

请你返回 nums 中 所有 整数对里，数位差之和。


*/

func sumDigitDifferences(nums []int) int64 {
	matrix := make([][10]int, len(strconv.Itoa(nums[0])))
	ans := int64(0)
	for k, x := range nums {

		for i := 0; x > 0; x /= 10 {
			d := x % 10
			ans += int64(k - matrix[i][d])
			matrix[i][d]++
			i++
		}
	}
	return ans
}
