package main

/**
题目：2425. 所有数对的异或和

给你两个下标从 0 开始的数组 nums1 和 nums2 ，
两个数组都只包含非负整数。请你求出另外一个数组 nums3 ，
包含 nums1 和 nums2 中 所有数对 的异或和（nums1 中每个整数都跟 nums2 中每个整数 恰好 匹配一次）。

请你返回 nums3 中所有整数的 异或和 。


*/

func xorAllNums(nums1 []int, nums2 []int) int {
	ans := 0
	if len(nums2)%2 > 0 {
		for _, x := range nums1 {
			ans ^= x
		}
	}
	if len(nums1)%2 > 0 {
		for _, x := range nums2 {
			ans ^= x
		}
	}
	return ans
}
