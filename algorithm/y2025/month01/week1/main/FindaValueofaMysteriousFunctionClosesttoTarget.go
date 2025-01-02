package main

import "math"

/**
 * 题目：1521. 找到最接近目标值的函数值
 *
 * Winston 构造了一个如上所示的函数 func 。他有一个整数数组 arr 和一个整数 target ，
 * 他想找到让 |func(arr, l, r) - target| 最小的 l 和 r 。
 *
 * 请你返回 |func(arr, l, r) - target| 的最小值。
 *
 * 请注意， func 的输入参数 l 和 r 需要满足 0 <= l, r < arr.length 。
 */

func closestToTarget(arr []int, target int) int {
	N := len(arr)
	ans := math.MaxInt
	for i := 0; i < N; i++ {
		x := arr[i]

		ans = min(ans, abs(x-target))
		for j := i - 1; j >= 0 && (arr[j]&x) != arr[j]; j++ {
			arr[j] &= x
			ans = min(ans, abs(arr[j]-target))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
