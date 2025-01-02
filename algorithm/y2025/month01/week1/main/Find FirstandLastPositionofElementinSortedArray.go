package main

/**
题目：34. 在排序数组中查找元素的第一个和最后一个位置
标签：二分查找

给你一个按照非递减顺序排列的整数数组 nums，和一个目标值 target。
请你找出给定目标值在数组中的开始位置和结束位置。

如果数组中不存在目标值 target，返回 [-1, -1]。

你必须设计并实现时间复杂度为 O(log n) 的算法解决此问题。


*/

func searchRange(nums []int, target int) []int {
	start := lowbound(nums, target)
	if start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := lowbound(nums, target+1) - 1
	return []int{start, end}
}

// 闭区间写法
func lowbound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1 // 闭区间 [left,right]
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid - 1 //闭区间 [left,mid-1]
		} else {
			left = mid + 1 //闭区间 [mid+1,right]
		}
	}
	return left
}
