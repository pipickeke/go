package main

func longestNiceSubarray(nums []int) int {

	ans := 0
	for i, or := range nums {
		j := i - 1
		for ; j >= 0 && (or&nums[j]) == 0; j-- {
			or |= nums[j]
		}
		ans = max(ans, i-j)
	}
	return ans
}

func longestNiceSubarray_2(nums []int) int {

	ans := 0
	left, or := 0, 0
	for right, x := range nums {
		for or&x > 0 {
			or ^= nums[left]
			left += 1
		}
		or |= x
		ans = max(ans, right-left+1)
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
